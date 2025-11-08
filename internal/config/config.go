package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	LocalDir   string `json:"localdir"`
	RemoteURL  string `json:"remoteurl"`
	RemotePath string `json:"remotepath"`
	User       string `json:"username"`
	Pass       string `json:"password"`
}

func (cfg *Config) Load(path string) error {

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &cfg) // преобразование json в структуру Config{}
	if err != nil {
		return err
	}

	return nil
}
