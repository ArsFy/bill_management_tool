package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func isJsonExist(filename string) bool {
	wd, _ := os.Getwd()
	if !isExists(path.Join(wd, "./data")) {
		os.Mkdir(path.Join(wd, "./data"), 0666)
	}
	return isExists(path.Join(wd, "./data/", filename))
}

func WriteFile(filPth, fileText string) {
	file, err := os.OpenFile(filPth, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("File Open", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(fileText)
	write.Flush()
}
