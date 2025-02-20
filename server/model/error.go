package model

import "errors"

var (
	ERROR_USER_QUIT = errors.New("用戶已退出")
	ERROR_USER_NOT_EXISTS = errors.New("用戶不存在")
	ERROR_USER_EXISTS = errors.New("用戶已存在")
	ERROR_USER_PWD = errors.New("密碼錯誤")

	ERROR_SERVER = errors.New("伺服器錯誤")
)
