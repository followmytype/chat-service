package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"client/model"
	"net"
)

func ReadConnMessage(conn net.Conn) (msg model.Message, err error) {
	buf := [8096]byte{}
	_, err = conn.Read(buf[:4]) // 這裡會阻塞等待對方傳送訊息
	if err != nil {
		fmt.Println("header讀取失敗：", err)
		return
	}
	// 收到的請求長度
	msgLen := binary.BigEndian.Uint32(buf[:4])

	n, err := conn.Read(buf[:msgLen])
	if n != int(msgLen) || err != nil {
		fmt.Println("body讀取失敗：", err)
		return
	}

	err = json.Unmarshal(buf[:msgLen], &msg)
	if err != nil {
		fmt.Println("json解析失敗：", err)
		return
	}

	return
}

func MakeConnMessage(msgType string, payload any) (msgJson []byte, err error) {
	msgData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}

	message := model.Message{}
	message.Type = msgType
	message.Data = string(msgData)

	msgJson, err = json.Marshal(message)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}

	return
}

func WriteConnMessage(conn net.Conn, data []byte) (err error) {
	// 發送資料到伺服器時，先傳送此次發送資料的長度
	// 因為長度為整數，而發送的型態為byte數組，所以要先將整數轉為[]byte
	// 事先準備一個長度為4的byte切片
	buf := [8096]byte{}
	msgLen := uint32(len(data))
	binary.BigEndian.PutUint32(buf[:4], msgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("header發送失敗", err)
		return
	}

	n, err = conn.Write(data)
	if n != int(msgLen) || err != nil {
		fmt.Println("body發送失敗", err)
		return
	}
	return
}
