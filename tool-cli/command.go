package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
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

func useObj(name string) {
	useRes, err := httpPost(config.Server+"/api/is_obj_exist", "name="+name)
	if err != nil {
		fmt.Printf("Server Err: %e", err)
		return
	}
	var useJson map[string]interface{}
	json.Unmarshal([]byte(useRes), &useJson)
	if useJson["isExist"].(bool) {
		use = name
		fmt.Printf("Select: %s", name)
	} else {
		fmt.Printf("Select Err: %s is not exist, use 'create obj_name'", name)
	}
}

func add(operator string, change int, comment string, timestamp int) {
	var thisTime string = fmt.Sprint(time.Now().Unix())
	if timestamp != -1 {
		thisTime = fmt.Sprint(timestamp)
	}

	addRes, err := httpPost(config.Server+"/api/add_record", fmt.Sprintf("name=%s&operator=%s&change=%s&comment=%s&time=%s", use, operator, fmt.Sprint(change), comment, thisTime))
	if err != nil {
		fmt.Printf("Server Err: %e", err)
		return
	}
	var addJson map[string]interface{}
	json.Unmarshal([]byte(addRes), &addJson)
	if addJson["status"].(float64) == 200 {
		fmt.Printf("Add successfully")
	} else {
		fmt.Printf("err: %s", addJson["err"].(string))
	}
}
