package net

import (
	"fmt"
	"net"
	"time"

	"github.com/hakuromi/webdav-sync/logger"
)

// Проверка интернет соединения
func CheckInternetConnection() error {
	_, err := net.DialTimeout("tcp", "8.8.8.8:53", 3*time.Second)
	if err != nil {
		logger.LogError("Internet connection error!", err)
		return err
	}
	return nil
}

func WaitForInternet() {
	for {
		err := CheckInternetConnection()
		if err == nil {
			return // если соединение есть, выходим из функции
		}
		logger.LogError("Internet is unavailable. Try synch again in 5 minutes.", nil)
		fmt.Println("Internet is unavailable. Try synch again in 5 minutes.")
		time.Sleep(5 * time.Minute)
	}
}
