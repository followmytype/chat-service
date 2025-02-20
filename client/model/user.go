package model

import "net"

const (
	MENU = iota
	ROOM
)

type User struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	Status     int    `json:"status"`
	Conn       net.Conn
	Layer      int
	ChatUserId string
}
