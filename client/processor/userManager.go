package processor

import (
	"fmt"
	"client/model"
	"net"
)

var self *model.User
var onlineUsers map[string]*model.User = make(map[string]*model.User)

func onlineSelf(userId, userName string, conn net.Conn) {
	self = &model.User{
		UserId:   userId,
		UserName: userName,
		Status:   model.ONLINE,
		Conn:     conn,
	}
}

func showOnlineUsers(startChat bool) {
	fmt.Println("當前在線用戶：")
	for userId, user := range onlineUsers {
		fmt.Printf("%s - %s\n", userId, model.StatusMap[user.Status])
	}
	if startChat {
		fmt.Println("請輸入用戶id進入聊天室，輸入0或空白退出")
		userId := ""
		fmt.Scanln(&userId)
		if userId == "0" || userId == "" {
			return
		}
		openRoom(userId)
	}
}

func updateUserStatus(userId string, status int) {
	user, ok := onlineUsers[userId]
	if !ok {
		onlineUsers[userId] = &model.User{UserId: userId, Status: status}
		return
	}
	user.Status = status
	switch status {
	case model.ONLINE:
		onlineUsers[userId] = user
	case model.BUSY:
		onlineUsers[userId] = user
	case model.OFFLINE:
		delete(onlineUsers, userId)
	}
}
