package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"server/model"
	"server/processor"
	"server/utils"
)

type MainProcessor struct {
	Conn net.Conn
	Up   *processor.UserProcessor
}

func (mp *MainProcessor) process() (err error) {
	for {
		msg, err := utils.ReadConnMessage(mp.Conn)
		if err != nil {
			if err == io.EOF {
				mp.connectionOver()
			} else {
				fmt.Println("讀取請求錯誤：", err)
			}
			return err
		}
		if err := mp.serverProcessMsg(&msg); err != nil {
			return err
		}
	}
}

func (mp *MainProcessor) serverProcessMsg(msg *model.Message) (err error) {
	switch msg.Type {
	case model.LOGIN_MSG_TYPE:
		mp.Up = &processor.UserProcessor{
			Conn: mp.Conn,
		}
		return mp.Up.ServerProcessLogin(msg)
	case model.REGISTER_MSG_TYPE:
		up := &processor.UserProcessor{
			Conn: mp.Conn,
		}
		return up.ServerProcessRegister(msg)
	case model.SMS_MSG_TYPE:
		if err = processor.SP.TransferClientMessage(msg); err != nil {
			fmt.Println("轉發訊息失敗", err)
		}
		return nil
	default:
		fmt.Println("request type error")
		return errors.New("request type error ")
	}
}

func (mp *MainProcessor) connectionOver() {
	if mp.Up != nil {
		mp.Up.Offline()
	}
}
