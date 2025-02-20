package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"server/model"
	"server/utils"

	"github.com/redis/go-redis/v9"
)

type UserRepository struct {
	rdb *redis.Client
}

var userRepository *UserRepository

func init() {
	userRepository = &UserRepository{
		utils.Rdb(),
	}
}

func User() *UserRepository {
	return userRepository
}

func (ur *UserRepository) GetById(userId string) (user model.User, err error) {
	ctx := context.Background()
	val, err := ur.rdb.HGet(ctx, "users", userId).Result()
	if err != nil {
		if err == redis.Nil {
			return user, model.ERROR_USER_NOT_EXISTS
		}
		return
	}
	err = json.Unmarshal([]byte(val), &user)

	return
}

func (ur *UserRepository) Create(user model.User) (err error) {
	ctx := context.Background()
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json marshal error")
		return err
	}
	err = ur.rdb.HSet(ctx, "users", user.UserId, string(userJson)).Err()
	if err != nil {
		fmt.Println("新增用戶失敗")
		return err
	}
	return
}
