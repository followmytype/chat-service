package main

import (
	"client/processor"
	"client/utils"
	"fmt"
	"os"
)



func main() {
	var loop = true
	fmt.Println("Server Address:", utils.GetServerAddr())

	fmt.Println("簡單聊天室")
	fmt.Println("\n1 登入")
	fmt.Println("2 註冊")
	fmt.Println("3 離開")
	fmt.Print("\n請輸入 1-3：")

	for loop {
		key, _ := utils.ScanNumber()
		switch key {
		case 1:
			fmt.Print("用戶id：")
			userId := utils.Scan()
			fmt.Print("密碼：")
			userPwd := utils.Scan()
			up := &processor.UserProcessor{}
			if err := up.Login(userId, userPwd); err != nil {
				return
			}
		case 2:
			fmt.Println("用戶id：")
			userId := utils.Scan()
			fmt.Println("密碼：")
			userPwd := utils.Scan()
			fmt.Println("用戶名：")
			userName := utils.Scan()
			up := &processor.UserProcessor{}
			if err := up.Register(userId, userPwd, userName); err != nil {
				return
			}
		case 3:
			os.Exit(0)
		default:
			fmt.Print("請重新輸入：")
		}
	}
}
