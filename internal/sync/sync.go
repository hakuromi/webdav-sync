package sync

import (
	"os"
	"path/filepath"
	"time"

	"github.com/studio-b12/gowebdav"
)

func Upload(client *gowebdav.Client, localpath, remotepath string) error {
	err := client.Remove(remotepath)
	if err != nil {
		return err
	}
	file, err := os.Open(localpath) // получение указателя на открытый файл
	if err != nil {
		return err
	}
	defer file.Close() // закрываем файл после выполнения функции

	err = client.WriteStream(remotepath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LatestFile(dir string) (string, error) {
	var latestPath string    // путь к последнему созданному файлу для синхронизации
	var latestTime time.Time // время создания файла

	file, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	for _, f := range file {
		//os.Stat получает информацию о файле, в info вернется структура os.FileInfo
		info, err := os.Stat(filepath.Join(dir, f.Name())) //формируем полный путь до файла для получения информации о нем
		if err != nil {
			continue
		}

		if info.ModTime().After(latestTime) {
			latestTime = info.ModTime()
			latestPath = filepath.Join(dir, f.Name())
		}
	}
	return latestPath, nil
}
