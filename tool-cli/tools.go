package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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

func httpPost(url, postinfo string) (string, error) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded", strings.NewReader(postinfo))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
