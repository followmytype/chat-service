package model

import (
)

type User struct {
	UserId   string `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
