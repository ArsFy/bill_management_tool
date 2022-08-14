package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var config map[string]interface{}

func init() {
	wd, _ := os.Getwd()
	configFile, err := ioutil.ReadFile(path.Join(wd, "config.json"))
	if err != nil {
		fmt.Println("Config error:", "please read README.md to configure 'config.json' in the current directory.")
		os.Exit(0)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("Config error:", "please read README.md to configure 'config.json' in the current directory.")
		os.Exit(0)
	}
}
