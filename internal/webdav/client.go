package webdav

import (
	"github.com/studio-b12/gowebdav"
)

func NewClient(url, user, pass string) (*gowebdav.Client, error) { // принимает данные из нашего конфига, возвращает клиента
	return gowebdav.NewClient(url, user, pass), nil // возвращается результат выполнения функции NewClient
}

func Connect(client *gowebdav.Client) error {
	return client.Connect()
}
