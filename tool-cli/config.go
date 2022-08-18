package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type ConfigType struct {
	Init   bool   `json:"init"`
	Server string `json:"server"`
	Last   string `json:"last"`
}

var config ConfigType

func init() {
	wd, _ := os.Getwd()
	configFile, err := ioutil.ReadFile(path.Join(wd, "config.json"))
	if err != nil {
		return
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("Config error:", "Formatting error, use default config")
		return
	}

	use = config.Last
}

func configUpdata() {
	jsonStr, _ := json.Marshal(config)
	wd, _ := os.Getwd()

	WriteFile(path.Join(wd, "config.json"), string(jsonStr))
}
