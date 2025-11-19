package logger

import (
	"log"
	"os"
	"path/filepath"
)

var logFile *os.File
var LogPath string

// initLogger настраивает стандартный пакет log: указывает файл вывода и формат меток.
func InitLogger() {
	exePath, err := os.Executable() // Получаем путь к исполняемому файлу
	if err != nil {
		log.Fatal(err)
	}

	// Формируем путь к файлу логов рядом с main.exe
	LogPath = filepath.Join(filepath.Dir(exePath), "log.txt")

	// Открываем (или создаём) файл в режиме дописать в конец.
	file, err := os.OpenFile(LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //os.OpenFile открывает или создает файл
	if err != nil {
		log.Fatal(err)
	}

	logFile = file
	// Указываем пакету log писать в этот файл
	log.SetOutput(file)

	// LstdFlags добавит дату и время к каждой записи
	log.SetFlags(log.LstdFlags)
}

func LogError(text string, err error) { // Пишет в лог-файл ошибку
	if err != nil {
		log.Printf("Error: %s. %v", text, err)
	}
}

func Log(text string) {
	log.Println(text)
}

func LogFatal(text string, err error) { //Пишет в лог-файл фатальную ошибку
	if err != nil {
		log.Fatalf("Fatal error: %s. %v", text, err)
	}
}

func Close() {
	if logFile != nil {
		logFile.Close()
	}
}
