package model

const (
	LOGIN_MSG_TYPE     = "LoginMsg"
	LOGIN_RES_MSG_TYPE = "LoginResMsg"

	REGISTER_MSG_TYPE     = "RegisterMsg"
	REGISTER_RES_MSG_TYPE = "RegisterResMsg"

	ONLINE_NOTIFY_MSG_TYPE = "OnlineNotifyMsg"

	SMS_MSG_TYPE            = "SmsMsg"
	SMS_MSG_COLLECTION_TYPE = "SmsMsgCollection"
)

const (
	ONLINE = iota
	OFFLINE
	BUSY
)

var StatusMap = map[int]string{
	ONLINE:  "online",
	OFFLINE: "offline",
	BUSY:    "busy",
}

// 資料傳送的格式
type Message struct {
	Type string // 資料類型
	Data string // 資料內容
}

// ---Message.Data內的格式---
// 登入資料
type LoginMsg struct {
	UserId   string
	UserPwd  string
	UserName string
}

// 登入回傳資料
type LoginResMsg struct {
	Code         int            // 登入結果狀態碼
	OnlineUserId map[string]int // 當前在線用戶id
	Error        string         // 登入錯誤訊息
}

// 註冊資料
type RegisterMsg struct {
	UserId   string
	UserPwd  string
	UserName string
}

// 註冊回傳資料
type RegisterResMsg struct {
	Code  int    // 註冊結果狀態碼
	Error string // 註冊錯誤訊息
}

// 用戶上線通知
type OnlineNotifyMsg struct {
	UserId string
	Status int
}

// 短訊資料
type SmsMsg struct {
	From    string
	To      string
	Content string
}

type SmsMsgCollection struct {
	SmsMsgs []SmsMsg
}
