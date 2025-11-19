package sync

import (
	"os"
	"path"
	"path/filepath"
	"sort"
	"time"

	"github.com/studio-b12/gowebdav"
)

type fileFromDir struct {
	Path    string
	ModTime time.Time
}

func Upload(client *gowebdav.Client, localpaths []string, baseremotepath string) error {
	filesForDelete, err := client.ReadDir(baseremotepath) // читаем папку на сервере
	if err != nil {
		return err
	}

	for _, del := range filesForDelete { // в цикле проходимся по всем файлам из папки на сервере
		deletePath := path.Join(baseremotepath, del.Name())
		err := client.Remove(deletePath)
		if err != nil {
			return err
		}
	}

	for _, lpath := range localpaths {
		//remotepath := fmt.Sprintf("%s_%d.rar", baseremotepath, i+1) // путь для загрузки файла на сервер

		remotepath := baseremotepath + "/" + filepath.Base(lpath)

		if _, err := client.Stat(remotepath); err == nil { // проверка наличия файла на сервере
			err = client.Remove(remotepath)
			if err != nil {
				return err
			}
		}

		file, err := os.Open(lpath) // получение указателя на открытый файл
		if err != nil {
			return err
		}

		defer file.Close() // закрываем файл после выполнения функции

		err = client.WriteStream(remotepath, file, 0644)
		if err != nil {
			return err
		}

	}

	return nil
}

func LatestFile(dir string, count int) ([]string, error) {
	var allFiles []fileFromDir
	var latestPaths []string // путь к последнему созданному файлу для синхронизации
	//var latestTime time.Time // время создания файла

	file, err := os.ReadDir(dir) // слайс всех папок внутри директории
	if err != nil {
		return nil, err
	}

	for _, f := range file {
		//os.Stat получает информацию о файле, в info вернется структура os.FileInfo
		info, err := os.Stat(filepath.Join(dir, f.Name())) //формируем полный путь до файла для получения информации о нем
		if err != nil {
			continue // если не смогли получить информацию о файле, перехлдим к следующему
		}

		allFiles = append(allFiles, fileFromDir{ // добавляем все файлы из директории в слайс структур
			Path:    filepath.Join(dir, f.Name()),
			ModTime: info.ModTime(),
		})
	}

	sort.Slice(allFiles, func(i, j int) bool {
		return allFiles[i].ModTime.After(allFiles[j].ModTime)
	})

	for i := 0; i < count; i++ {
		latestPaths = append(latestPaths, allFiles[i].Path)
	}

	return latestPaths, nil
}
