package utils

import "os"

var serverAddr string

func init() {
	if len(os.Args) > 1 {
		serverAddr = os.Args[1]
        return 
	}
    serverAddr = GetEnv("SERVER_ADDR", "localhost:8889")
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func GetServerAddr() string {
    return serverAddr
}
