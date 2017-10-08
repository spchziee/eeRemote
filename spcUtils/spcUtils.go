package spcUtils

import (
	"bytes"
	"compress/zlib"
	"io"
)

func Byte2StringWithEnd(data []byte) string {
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			return string(data[0:i])
		}
	}
	return string(data)
}

//进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}
