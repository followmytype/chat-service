package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"server/model"
	"server/repository"
	"server/utils"
)

type UserProcessor struct {
	Conn   net.Conn
	UserId string
	Status int
}

func (up *UserProcessor) ServerProcessLogin(msg *model.Message) (err error) {
	var loginMsg model.LoginMsg
	if err = json.Unmarshal([]byte(msg.Data), &loginMsg); err != nil {
		fmt.Println("json解構失敗")
		return err
	}
	loginResMsg := model.LoginResMsg{}
	user, err := repository.User().GetById(loginMsg.UserId)
	if err == model.ERROR_USER_NOT_EXISTS {
		loginResMsg.Code = 404
		loginResMsg.Error = err.Error()
	} else {
		if user.UserPwd == loginMsg.UserPwd {
			loginResMsg.Code = 200
			up.UserId = loginMsg.UserId
			up.Status = model.ONLINE
			userMgr.AddOnlineUser(up)
			loginResMsg.OnlineUserId = userMgr.GetAllOnlineUserIdStatusMap()
			go up.ServerNotifyOnlineUser(loginMsg.UserId, model.ONLINE)
			fmt.Printf("%s上線\n", loginMsg.UserId)
		} else {
			loginResMsg.Code = 403
			loginResMsg.Error = model.ERROR_USER_PWD.Error()
		}
	}
	responseJson, err := utils.MakeConnMessage(model.LOGIN_RES_MSG_TYPE, loginResMsg)
	if err != nil {
		return err
	}
	err = utils.WriteConnMessage(up.Conn, responseJson)
	if err != nil {
		return err
	}
	if loginResMsg.Code == 200 {
		SP.SendLeavedMsg(loginMsg.UserId)
	}
	return
}

func (up *UserProcessor) ServerProcessRegister(msg *model.Message) (err error) {
	var registerMsg model.RegisterMsg
	if err = json.Unmarshal([]byte(msg.Data), &registerMsg); err != nil {
		fmt.Println("json解構失敗")
		return err
	}
	registerResMsg := model.RegisterResMsg{Code: 200, Error: ""}
	user, err := repository.User().GetById(registerMsg.UserId)
	if err != nil {
		if err == model.ERROR_USER_NOT_EXISTS {
			user.UserId = registerMsg.UserId
			user.UserPwd = registerMsg.UserPwd
			user.UserName = registerMsg.UserName
			if err = repository.User().Create(user); err != nil {
				registerResMsg.Code = 500
				registerResMsg.Error = err.Error()
			}
			fmt.Printf("%s註冊\n", registerMsg.UserId)
		} else {
			registerResMsg.Code = 500
			registerResMsg.Error = err.Error()
		}
	} else {
		registerResMsg.Code = 400
		registerResMsg.Error = model.ERROR_USER_EXISTS.Error()
	}
	msgJson, err := utils.MakeConnMessage(model.REGISTER_RES_MSG_TYPE, registerResMsg)
	if err != nil {
		return err
	}
	return utils.WriteConnMessage(up.Conn, msgJson)
}

func (up *UserProcessor) ServerNotifyOnlineUser(userId string, status int) (err error) {
	notify := model.OnlineNotifyMsg{UserId: userId, Status: status}
	msg, err := utils.MakeConnMessage(model.ONLINE_NOTIFY_MSG_TYPE, notify)
	if err != nil {
		return err
	}
	for onlineUserId, userProcessor := range userMgr.onlineUsers {
		if onlineUserId == userId {
			continue
		}
		go utils.WriteConnMessage(userProcessor.Conn, msg)
	}
	return
}

func (up *UserProcessor) Offline() {
	userMgr.DeleteOnlineUser(up)
	up.ServerNotifyOnlineUser(up.UserId, model.OFFLINE)
	fmt.Printf("%s已離線\n", up.UserId)
}
