package main

import (
	"encoding/json"
	"fmt"
)

var v string = "0.1"

func main() {
	if !config.Init {
		fmt.Printf("Server (e.g. 'https://example.com'): ")
		fmt.Scan(&config.Server)
		testRes, err := httpPost(config.Server+"/api/status", "")
		if err != nil {
			fmt.Println("Server Err:", err)
		}
		var testResJson map[string]interface{}
		err = json.Unmarshal([]byte(testRes), &testResJson)
		if err != nil {
			fmt.Println("Server Err:", err)
		}
		if testResJson["status"].(float64) == 200 {
			config.Init = true
			configUpdata()
			fmt.Printf("Connection successful: Server version: %s, pls rerun the tool-cli.", testResJson["version"])
		} else {
			fmt.Println("Server Err:", "Unknown error,", "err", testResJson["status"])
		}
		return
	}

	fmt.Printf("BMTool tool-cli v%s, try 'help' for more information.", v)
	for {
		fmt.Printf("\nBMTool> ")
		var command string = ""
		var value string = ""
		fmt.Scanln(&command, &value)

		switch command {
		case "help":
			fmt.Printf("help")
		case "list":
			list()
		case "create":
			if value != "" {
				createObj(value)
			} else {
				fmt.Printf("missing name: 'create xxx'")
			}
		default:
			fmt.Printf("'%s' not found", command)
		}
	}
}
