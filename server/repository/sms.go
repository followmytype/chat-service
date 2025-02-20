package repository

import (
	"context"
	"encoding/json"
	"server/model"
	"server/utils"

	"github.com/redis/go-redis/v9"
)

type SmsRepository struct {
	rdb *redis.Client
}

var smsRepository *SmsRepository

func init() {
	smsRepository = &SmsRepository{
		utils.Rdb(),
	}
}

func Sms() *SmsRepository {
	return smsRepository
}

func (sr *SmsRepository) SaveSms(sms model.SmsMsg) (err error) {
	ctx := context.Background()
	smsJson, err := json.Marshal(sms)
	_, err = sr.rdb.RPush(ctx, sms.To, string(smsJson)).Result()

	return
}

func (sr *SmsRepository) GetUserSms(userId string) (smsCollection model.SmsMsgCollection, err error) {
	ctx := context.Background()
	values, err := sr.rdb.LRange(ctx, userId, 0, -1).Result()
	if err != nil {
		return smsCollection, err
	}
	for _, v := range values {
		sms := model.SmsMsg{}
		json.Unmarshal([]byte(v), &sms)
		smsCollection.SmsMsgs = append(smsCollection.SmsMsgs, sms)
	}
	sr.rdb.LTrim(ctx, userId, 1, 0)
	return
}
