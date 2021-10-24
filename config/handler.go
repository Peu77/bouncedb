package config

import (
	"bouncedb/file"
	"crypto/rand"
	"fmt"
	"github.com/goccy/go-json"
)

type Config struct {
}

var CurrentConfig = Config{}

const configFile = "files/config.conf"

var Token []byte

func InitConfig() {
	Token = make([]byte, 32)
	rand.Read(Token)

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
