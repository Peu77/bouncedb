package config

import (
	"bouncedb/file"
	"fmt"
	"github.com/goccy/go-json"
)

type Config struct {
	SecretKey string
}

var CurrentConfig = Config{SecretKey: "secretKey1234!"}

const configFile = "/files/config.conf"

func InitConfig() {
	if file.ExistFile(configFile) {
		jsonString := file.ReadFile(configFile)
		err := json.Unmarshal([]byte(jsonString), &CurrentConfig)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		marshal, err := json.MarshalIndent(CurrentConfig, "", "  ")
		if err != nil {
			return
		}
		file.WriteInFile(configFile, string(marshal))
	}
}
