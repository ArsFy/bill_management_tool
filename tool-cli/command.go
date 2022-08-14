package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func list() {
	list, err := httpPost(config.Server+"/api/obj_list", "")
	if err != nil {
		fmt.Printf("Server Err: %e", err)
		return
	}
	var listJson map[string]interface{}
	json.Unmarshal([]byte(list), &listJson)

	if listJson["count"].(float64) != 0 {
		var liststr []string
		for _, j := range listJson["list"].([]interface{}) {
			liststr = append(liststr, j.(string))
		}
		fmt.Printf("Use 'use obj_name' to select a obj \n%s", strings.Join(liststr, "  "))
	} else {
		fmt.Printf("obj list is empty, use command 'create obj_name'")
	}
}

func createObj(name string) {
	create, err := httpPost(config.Server+"/api/create_obj", "name="+name)
	if err != nil {
		fmt.Printf("Server Err: %e", err)
		return
	}
	var createJson map[string]interface{}
	json.Unmarshal([]byte(create), &createJson)
	if createJson["status"].(float64) == 200 {
		switch createJson["info"].(string) {
		case "ok":
			fmt.Printf("Created successfully: %s", name)
		case "isExist":
			fmt.Printf("Created Error: %s is Exist", name)
		}
	}
}
