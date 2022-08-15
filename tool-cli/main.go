package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var v string = "0.1"
var use string = ""

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
		var value []string = make([]string, 4)
		fmt.Scanln(&command, &value[0], &value[1], &value[2], &value[3])

		switch command {
		case "help":
			fmt.Printf("help")
		case "list":
			list()
		case "create":
			if value[0] != "" {
				createObj(value[0])
			} else {
				fmt.Printf("missing name: 'create obj_name'")
			}
		case "use":
			if value[0] != "" {
				useObj(value[0])
			} else {
				fmt.Printf("input value missing: 'use obj_name'")
			}
		case "add":
			if value[0] != "" && value[1] != "" && value[2] != "" {
				if use != "" {
					changeNumber, _ := strconv.Atoi(value[1])
					var timestamp int = -1
					if value[3] != "" {
						timestamp, _ = strconv.Atoi(value[3])
					}
					add(value[0], changeNumber, value[2], timestamp)
				} else {
					fmt.Printf("must be selected an obj to use: 'use obj_name'")
				}
			} else {
				fmt.Printf("input value missing: 'add [Operator Name] [+123 or -123] [what is it used for] [Timestamp (optional)]'")
			}
		default:
			fmt.Printf("'%s' not found", command)
		}
	}
}
