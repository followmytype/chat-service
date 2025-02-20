package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"client/model"
	"client/utils"
	"net"
)

type UserProcessor struct {
}

func (up *UserProcessor) Login(userId, userPwd string) (err error) {
	conn, err := net.Dial("tcp", utils.GetServerAddr())
	if err != nil {
		fmt.Println("連線失敗：", err)
		return
	}
	defer conn.Close()

	loginMsg := model.LoginMsg{
		UserId:  userId,
		UserPwd: userPwd,
	}
	msgBytes, err := utils.MakeConnMessage(model.LOGIN_MSG_TYPE, loginMsg)
	if err != nil {
		return err
	}
	err = utils.WriteConnMessage(conn, msgBytes)
	if err != nil {
		return err
	}
	msg, err := utils.ReadConnMessage(conn)
	if err != nil {
		return err
	}
	var loginResMsg model.LoginResMsg
	if err = json.Unmarshal([]byte(msg.Data), &loginResMsg); err != nil {
		fmt.Println("payload解構失敗")
		return err
	}
	if loginResMsg.Code == 200 {
		go serverMonitor(conn) // 監聽伺服器端訊息
		onlineSelf(loginMsg.UserId, loginMsg.UserName, conn)
		fmt.Println(userId + "登入成功")
		for userId, status := range loginResMsg.OnlineUserId {
			if userId != loginMsg.UserId {
				updateUserStatus(userId, status)
			}
		}
		showOnlineUsers(false)
		StartService()
	}
	fmt.Println(loginResMsg.Error)

	return errors.New(loginResMsg.Error)
}

func (up *UserProcessor) Register(userId, userPwd, userName string) (err error) {
	conn, err := net.Dial("tcp", utils.GetServerAddr())
	if err != nil {
		fmt.Println("net dial error:", err)
		return
	}
	defer conn.Close()

	registerMsg := model.RegisterMsg{
		UserId:   userId,
		UserPwd:  userPwd,
		UserName: userName,
	}
	msgBytes, err := utils.MakeConnMessage(model.REGISTER_MSG_TYPE, registerMsg)
	if err != nil {
		return err
	}

	err = utils.WriteConnMessage(conn, msgBytes)
	if err != nil {
		return err
	}
	msg, err := utils.ReadConnMessage(conn)
	if err != nil {
		return err
	}
	var registerResMsg model.RegisterResMsg
	if err = json.Unmarshal([]byte(msg.Data), &registerResMsg); err != nil {
		fmt.Println("payload解構失敗")
		return err
	}
	if registerResMsg.Code == 200 {
		fmt.Println("註冊成功，請重新登入")
		return
	}
	fmt.Println(registerResMsg.Error)

	return errors.New(registerResMsg.Error)
}
