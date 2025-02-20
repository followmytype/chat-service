package processor

type UserMgr struct {
	onlineUsers map[string]*UserProcessor
}

var userMgr *UserMgr

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[string]*UserProcessor, 1024),
	}
}

func (um *UserMgr) AddOnlineUser(up *UserProcessor) {
	um.onlineUsers[up.UserId] = up
}

func (um *UserMgr) DeleteOnlineUser(up *UserProcessor) {
	delete(um.onlineUsers, up.UserId)
}

func (um *UserMgr) GetOnlineUserById(userId string) *UserProcessor {
	if up, ok := um.onlineUsers[userId]; ok {
		return up
	}
	return nil
}

func (um *UserMgr) GetAllOnlineUser() map[string]*UserProcessor {
	return um.onlineUsers
}

func (um *UserMgr) GetAllOnlineUserIdStatusMap() map[string]int {
	userIdStatusMap := make(map[string]int)
	for userId, up := range um.onlineUsers {
		userIdStatusMap[userId] = up.Status
	}
	return userIdStatusMap
}
