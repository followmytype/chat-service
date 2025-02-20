package processor

import (
	"client/model"
	"client/utils"
	"fmt"
	"sync"
)

var smsHistoryLock sync.Mutex
var smsHistory map[string][]model.SmsMsg

func init() {
	smsHistory = make(map[string][]model.SmsMsg)
}

func openRoom(userId string) {
	printSmsHistory(userId)
	self.ChatUserId = userId
	self.Status = model.BUSY
	// TODO: notify server change status
	fmt.Printf("開始與%s對話，輸入QUIT退出對話\n", userId)
	content := ""
	for {
		content = utils.Scan()
		if content == "QUIT" {
			self.ChatUserId = ""
			self.Status = model.ONLINE
			break
		}
		if err := SP.SendSms(userId, content); err != nil {
			fmt.Println("訊息發送錯誤")
		}
	}
}

func recordSmsHistory(sms model.SmsMsg) {
	smsHistoryLock.Lock()
	if sms.From == self.UserId {
		smsHistory[sms.To] = append(smsHistory[sms.To], sms)
	} else {
		smsHistory[sms.From] = append(smsHistory[sms.From], sms)
	}
	smsHistoryLock.Unlock()
}

func printSmsHistory(userId string) {
	smsHistoryLock.Lock()
	for _, sms := range smsHistory[userId] {
		if sms.From == userId {
			fmt.Printf("%s: %s\n", sms.From, sms.Content)
		} else {
			fmt.Println(sms.Content)
		}
	}
	smsHistoryLock.Unlock()
}

func showChatRoom() {
	smsHistoryLock.Lock()
	for room, msgs := range smsHistory {
		fmt.Printf("%s - %s\n", room, msgs[len(msgs)-1].Content)
	}
	smsHistoryLock.Unlock()
}
