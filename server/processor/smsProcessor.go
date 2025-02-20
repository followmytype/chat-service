package processor

import (
	"encoding/json"
	"server/model"
	"server/repository"
	"server/utils"
)

type SmsProcessor struct {
}

var SP *SmsProcessor

func init() {
	SP = &SmsProcessor{}
}

func (sp *SmsProcessor) TransferClientMessage(msg *model.Message) (err error) {
	var smsMsg model.SmsMsg
	if err = json.Unmarshal([]byte(msg.Data), &smsMsg); err != nil {
		return err
	}
	msgJson, err := utils.MakeConnMessage(model.SMS_MSG_TYPE, smsMsg)
	if err != nil {
		return err
	}
	if smsMsg.To == "" {
		for _, up := range userMgr.GetAllOnlineUser() {
			if up.UserId != smsMsg.From {
				go utils.WriteConnMessage(up.Conn, msgJson)
			}
		}
		return nil
	}
	up, ok := userMgr.onlineUsers[smsMsg.To]
	if !ok {
		_, err := repository.User().GetById(smsMsg.To)
		if err != nil {
			return err
		}
		return repository.Sms().SaveSms(smsMsg)
	}
	return utils.WriteConnMessage(up.Conn, msgJson)
}

func (sp *SmsProcessor) SendLeavedMsg(userId string) (err error) {
	msgCollection, err := repository.Sms().GetUserSms(userId)
	if err != nil || len(msgCollection.SmsMsgs) == 0 {
		return err
	}
	json, err := utils.MakeConnMessage(model.SMS_MSG_COLLECTION_TYPE, msgCollection)
	if err != nil {
		return err
	}
	err = utils.WriteConnMessage(userMgr.onlineUsers[userId].Conn, json)
	if err != nil {
		return err
	}
	return
}
