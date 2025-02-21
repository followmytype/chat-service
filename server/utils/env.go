package utils

import (
	"os"
	"strconv"
)

var serverAddr string
var port string

func init() {
	if len(os.Args) > 1 && isNumeric(os.Args[1]) {
		port = os.Args[1]
		serverAddr = "0.0.0.0:" + port
		return
	}
	port = GetEnv("PORT", "8889")
	serverAddr = "0.0.0.0:" + port
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetPort() string {
	return port
}

func GetServerAddr() string {
	return serverAddr
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
