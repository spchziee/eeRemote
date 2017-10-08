package fileTrans

import (
	"bytes"
	"eePACS/spcUtils"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
)

/*
DWORD dwCmdType;
	char szFilePath[_MAX_PATH];
	DWORD dwFileLen;
	DWORD dwTransLen;
	DWORD dwExistFileLen;
*/
type CMD struct {
	Type         int32
	FilePath     [256]byte
	FileLen      int32
	TransLen     int32
	ExistFileLen int32
}

var len_of_CMD int

func init() {
	len_of_CMD = spcUtils.SizeStruct(new(CMD))
}

func (cmd *CMD) Prase(data []byte) (int, error) {
	if len_of_CMD > len(data) {
		return 0, errors.New("no enough data")
	}

	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &cmd.Type)
	binary.Read(buf, binary.LittleEndian, &cmd.FilePath)
	binary.Read(buf, binary.LittleEndian, &cmd.FileLen)
	binary.Read(buf, binary.LittleEndian, &cmd.TransLen)
	binary.Read(buf, binary.LittleEndian, &cmd.ExistFileLen)

	fmt.Printf("%+v\n", cmd)

	return len_of_CMD, nil
}

func handleFileTrans(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, len_of_CMD)
	nRead, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var cmd CMD
	_, err = cmd.Prase(buffer[0:nRead])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	filename := spcUtils.Byte2StringWithEnd(cmd.FilePath[0:])
	fmt.Println(filename)
	buffer = make([]byte, cmd.TransLen)
	totalRead := 0
	for i := 0; i < int(cmd.TransLen)/1024+1; i++ {
		tmp := make([]byte, 1024)
		nRead, err = conn.Read(tmp)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		copy(buffer[totalRead:], tmp)
		totalRead += nRead
		fmt.Println("读取到 ", nRead)
	}

	if int32(totalRead) != cmd.TransLen {
		fmt.Println("收到数据长度不一致 ", int32(totalRead), cmd.TransLen)
		return
	}

	unCompressData := spcUtils.DoZlibUnCompress(buffer)

	err = ioutil.WriteFile(filename, unCompressData, 666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
