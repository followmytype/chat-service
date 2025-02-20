package main

import (
	"fmt"
	"io"
	"net"
	"server/utils"
)

func main() {
	listen, err := net.Listen("tcp", utils.GetServerAddr())
	if err != nil {
		fmt.Println("服務啟動失敗：", err)
	}
	defer listen.Close()
	fmt.Println("服務啟動..")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接收客戶請求錯誤：", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	processor := &MainProcessor{
		Conn: conn,
	}
	if err := processor.process(); err != nil {
		if err != io.EOF {
			fmt.Println("通訊錯誤", err)
		}
		return
	}
}
