package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hakuromi/webdav-sync/internal/config"
	"github.com/hakuromi/webdav-sync/internal/net"
	"github.com/hakuromi/webdav-sync/internal/sync"
	"github.com/hakuromi/webdav-sync/internal/webdav"
	"github.com/hakuromi/webdav-sync/logger"
)

func main() {
	logger.InitLogger()                                                 // запуск логгера
	defer logger.Close()                                                // после завершения работы программы логгер закроет файл
	fmt.Println("------Log file created at:", logger.LogPath, "------") /////////////////
	time.Sleep(2 * time.Second)

	net.WaitForInternet()                                             // проверка соединения с интернетом
	fmt.Println("------Internet connection checked. All right------") ///////////////

	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		fmt.Print("Path to json configuration file: ") // путь к json
		reader := bufio.NewReader(os.Stdin)            // чтение config.json файла
		input, err := reader.ReadString('\n')
		if err != nil {
			logger.LogError("Failed to read config file path", err)
		}
		path = strings.TrimSpace(input) // trimspace убирает лишние пробелы с начала и с конца строки
	}
	conf := config.Config{} // создаем экземпляр структуры конфигурации
	err := conf.Load(path)
	if err != nil {
		logger.LogFatal("Failed to load config!", err)
	}
	fmt.Println("------Config loaded successfully!------") /////////////////
	fmt.Println("------Using config file:", path, "------")
	time.Sleep(2 * time.Second) ////////////////////

	// Успешная загрузка конфига
	log.Printf("Configuration loaded:\n WebDAV URL: %s\n Local directory: %s\n Serv directory: %s\n", conf.RemoteURL, conf.LocalDir, conf.RemotePath)

	//Создаем клиента с помощью импортированной библиотеки gowebdav
	client, err := webdav.NewClient(conf.RemoteURL, conf.User, conf.Pass) // url сервера/облака, логин, пароль
	if err != nil {
		logger.LogFatal("Failed to load client.", err)
	}

	err = webdav.Connect(client) // соединение с сервером
	if err != nil {
		logger.LogFatal("Connection failed.", err)
	}
	fmt.Println("------Connection to the server success!------") ///////////////////////
	time.Sleep(2 * time.Second)                                  /////////////////////

	latestFile, err := sync.LatestFile(conf.LocalDir)
	if err != nil {
		logger.LogError("Failed to find latest file!", err)
	}
	fmt.Println("------Upload starting...------")

	err = sync.Upload(client, latestFile, conf.RemotePath) // отправка файла на сервер
	if err != nil {
		logger.LogFatal("Backup failed.", err)
	}
	fmt.Println("------Upload successful!------") ////////////////////////
	time.Sleep(2 * time.Second)                   ///////////////////////
}
