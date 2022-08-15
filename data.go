package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type DataType struct {
	Time    int    `json:"time"`
	User    string `json:"user"`
	Change  int    `json:"change"`
	Comment string `json:"comment"`
}

type DataListType map[string]map[string]map[string][]DataType

func readData(name string) (DataListType, error) {
	var data DataListType
	wd, _ := os.Getwd()
	file, err := ioutil.ReadFile(path.Join(wd, "/data/", name+".json"))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func updataData(name string, content DataListType) {
	wd, _ := os.Getwd()
	jsonStr, err := json.Marshal(content)
	if err != nil {
		fmt.Println("UpData Err:", err)
		return
	}
	WriteFile(path.Join(wd, "/data/", name+".json"), string(jsonStr))
}
