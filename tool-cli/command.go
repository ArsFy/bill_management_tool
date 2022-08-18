package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type DataType struct {
	Time    int    `json:"time"`
	User    string `json:"user"`
	Change  int    `json:"change"`
	Comment string `json:"comment"`
}

func list() {
	list, err := httpPost(config.Server+"/api/obj_list", "")
	if err != nil {
		fmt.Printf("Server Err: %s", config.Server)
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
		fmt.Printf("Server Err: %s", config.Server)
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
		fmt.Printf("Server Err: %s", config.Server)
		return
	}
	var useJson map[string]interface{}
	json.Unmarshal([]byte(useRes), &useJson)
	if useJson["isExist"].(bool) {
		use = name
		config.Last = name
		configUpdata()
		fmt.Printf("Select: %s", name)
	} else {
		fmt.Printf("Select Err: %s is not exist, use 'create obj_name'", name)
	}
}

func add(operator string, change float64, comment string, timestamp int) {
	var thisTime string = fmt.Sprint(time.Now().Unix())
	if timestamp != -1 {
		thisTime = fmt.Sprint(timestamp)
	}

	addRes, err := httpPost(config.Server+"/api/add_record", fmt.Sprintf("name=%s&operator=%s&change=%s&comment=%s&time=%s", use, operator, fmt.Sprint(change*100), comment, thisTime))
	if err != nil {
		fmt.Printf("Server Err: %s", config.Server)
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

func info(user string) {
	infoRes, err := httpPost(config.Server+"/api/get_record", fmt.Sprintf("name=%s&user=%s", use, user))
	if err != nil {
		fmt.Printf("Server Err: %s", config.Server)
		return
	}

	var infoJson struct {
		Status int                   `json:"status"`
		Data   map[string][]DataType `json:"data"`
		Am     int                   `json:"am"`
		In     int                   `json:"in"`
		Out    int                   `json:"out"`
		Lm     int                   `json:"lm"`
	}
	json.Unmarshal([]byte(infoRes), &infoJson)

	if infoJson.Status == 200 {
		lm := float64(infoJson.Lm)
		var lms string
		if lm != 0.0 {
			lms = fmt.Sprintf("%.2f", ((float64(infoJson.Am) - lm) / lm))
		} else {
			lms = fmt.Sprintf("%.2f", float64(infoJson.Am))
		}

		maxname := 0
		maxchange := 0
		maxcomment := 0

		m := infoJson.Data

		for _, j := range m {
			for _, jj := range j {
				change := fmt.Sprintf("%.2f", float64(jj.Change)/100)
				if maxname < len(jj.User) {
					maxname = len(jj.User)
				}
				if maxchange < len(change) {
					maxchange = len(change)
				}
				if maxcomment < len(jj.Comment) {
					maxcomment = len(jj.Comment)
				}
			}
		}

		if maxname-4 < 0 {
			maxname = 4
		}
		if maxchange-6 < 0 {
			maxchange = 6
		}
		if maxcomment-7 < 0 {
			maxcomment = 7
		}

		text := ""
		for _, j := range m {
			for _, jj := range j {
				thisTime := time.Unix(int64(jj.Time), 0).Format("2006-01-02 15:04")
				change := fmt.Sprintf("%.2f", float64(jj.Change)/100)
				text = fmt.Sprintf("\n| %s | %s | %s | %s |", thisTime, jj.User+strings.Repeat(" ", maxname-len(jj.User)), change+strings.Repeat(" ", maxchange-len(change)), jj.Comment+strings.Repeat(" ", maxcomment-len(jj.Comment))) + text
			}
		}

		top := fmt.Sprintf("| Time             | Name%s | Change%s | Comment%s |", strings.Repeat(" ", maxname-4), strings.Repeat(" ", maxchange-6), strings.Repeat(" ", maxcomment-7))
		text = "+" + strings.Repeat("-", len(top)-2) + "+\n" + top + "\n+" + strings.Repeat("-", len(top)-2) + "+" + text + "\n+" + strings.Repeat("-", len(top)-2) + "+"

		fmt.Printf(
			"Time: %s\nthis month: %s | income: %s | expenses: %s | last month: %s\n\n%s",
			time.Now().Format("2006-01-02 15:04"),
			fmt.Sprintf("%.2f", float64(infoJson.Am)/100),
			fmt.Sprintf("%.2f", float64(infoJson.In)/100),
			fmt.Sprintf("%.2f", float64(infoJson.Out)/100),
			lms+"%",
			text,
		)
	} else {
		fmt.Println("err", infoJson)
	}
}

func year(user string) {
	infoRes, err := httpPost(config.Server+"/api/get_record_year", fmt.Sprintf("name=%s&user=%s", use, user))
	if err != nil {
		fmt.Printf("Server Err: %s", config.Server)
		return
	}

	var infoJson struct {
		Status int                              `json:"status"`
		Data   map[string]map[string][]DataType `json:"data"`
		Am     int                              `json:"am"`
		In     int                              `json:"in"`
		Out    int                              `json:"out"`
	}
	json.Unmarshal([]byte(infoRes), &infoJson)

	if infoJson.Status == 200 {
		maxname := 0
		maxchange := 0
		maxcomment := 0

		m := infoJson.Data

		for _, j := range m {
			for _, jj := range j {
				for _, jjj := range jj {
					change := fmt.Sprintf("%.2f", float64(jjj.Change)/100)
					if maxname < len(jjj.User) {
						maxname = len(jjj.User)
					}
					if maxchange < len(change) {
						maxchange = len(change)
					}
					if maxcomment < len(jjj.Comment) {
						maxcomment = len(jjj.Comment)
					}
				}
			}
		}

		if maxname-4 < 0 {
			maxname = 4
		}
		if maxchange-6 < 0 {
			maxchange = 6
		}
		if maxcomment-7 < 0 {
			maxcomment = 7
		}

		text := ""
		for _, j := range m {
			for _, jj := range j {
				for _, jjj := range jj {
					thisTime := time.Unix(int64(jjj.Time), 0).Format("2006-01-02 15:04")
					change := fmt.Sprintf("%.2f", float64(jjj.Change)/100)
					text = fmt.Sprintf("\n| %s | %s | %s | %s |", thisTime, jjj.User+strings.Repeat(" ", maxname-len(jjj.User)), change+strings.Repeat(" ", maxchange-len(change)), jjj.Comment+strings.Repeat(" ", maxcomment-len(jjj.Comment))) + text
				}
			}
		}

		top := fmt.Sprintf("| Time             | Name%s | Change%s | Comment%s |", strings.Repeat(" ", maxname-4), strings.Repeat(" ", maxchange-6), strings.Repeat(" ", maxcomment-7))
		text = "+" + strings.Repeat("-", len(top)-2) + "+\n" + top + "\n+" + strings.Repeat("-", len(top)-2) + "+" + text + "\n+" + strings.Repeat("-", len(top)-2) + "+"

		fmt.Printf(
			"Time: %s\nthis year: %s | income: %s | expenses: %s\n\n%s",
			time.Now().Format("2006-01-02 15:04"),
			fmt.Sprintf("%.2f", float64(infoJson.Am)/100),
			fmt.Sprintf("%.2f", float64(infoJson.In)/100),
			fmt.Sprintf("%.2f", float64(infoJson.Out)/100),
			text,
		)
	} else {
		fmt.Println("err", infoJson)
	}
}
