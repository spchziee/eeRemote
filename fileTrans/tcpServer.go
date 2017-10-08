package fileTrans

import (
	"fmt"
	"net"
)

func Start(addr string) {
	//侦听
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic("ListenTCP error")
	}
	//记得关闭
	defer listen.Close()

	fmt.Println("start server successful......")
	//开始接收请求
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("ListenTCP Accept error")
			continue
		}
		fmt.Println("accept tcp client ", conn.RemoteAddr())
		// 每次建立一个连接就放到单独的协程内做处理
		go handleFileTrans(conn)
	}
}
