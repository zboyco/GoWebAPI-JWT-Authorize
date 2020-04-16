package utils

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	Mode     string `json:"mode"`
	Database struct {
		Host   string `json:"host"`
		Port   uint16 `json:"port"`
		User   string `json:"user"`
		Pwd    string `json:"pwd"`
		DbName string `json:"db_name"`
	} `json:"database"`
	TokenExpires int `json:"token_expires"`
}

var Config Configuration

// ConfigInit 初始化配置文件
func ConfigInit(configPathSlice ...string) error {
	configPath := "config.json"
	if len(configPathSlice) == 1 {
		configPath = configPathSlice[0]
	}
	jsonByte, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonByte, &Config)
	if err != nil {
		return err
	}

	return nil
}
