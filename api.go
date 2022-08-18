package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func objList(c *gin.Context) {
	wd, _ := os.Getwd()
	if !isExists(path.Join(wd, "./data")) {
		os.Mkdir(path.Join(wd, "./data"), 0666)
		c.JSON(200, gin.H{"status": 200, "list": []string{}, "count": 0})
		return
	}
	files, _ := ioutil.ReadDir(path.Join(wd, "./data"))
	var objlist []string
	for _, f := range files {
		objlist = append(objlist, strings.Replace(f.Name(), ".json", "", 1))
	}

	c.JSON(200, gin.H{"status": 200, "list": objlist, "count": len(objlist)})
}

func createObj(c *gin.Context) {
	wd, _ := os.Getwd()
	name := c.PostForm("name")
	if name == "" {
		c.JSON(200, gin.H{"status": 500, "err": "empty name"})
		return
	}
	if isJsonExist(name + ".json") {
		c.JSON(200, gin.H{"status": 200, "info": "isExist"})
	} else {
		WriteFile(path.Join(wd, "./data/", name+".json"), "{}")
		c.JSON(200, gin.H{"status": 200, "info": "ok"})
	}
}

func isObjExist(c *gin.Context) {
	name := c.PostForm("name")
	c.JSON(200, gin.H{"status": 200, "isExist": isJsonExist(name + ".json")})
}

func addRecord(c *gin.Context) {
	name := c.PostForm("name")
	user := c.PostForm("operator")
	change := c.PostForm("change")
	changeNumber, _ := strconv.Atoi(change)
	comment := c.PostForm("comment")
	timeNumber, _ := strconv.ParseInt(c.PostForm("time"), 10, 64)

	data, err := readData(name)
	if err != nil {
		c.JSON(200, gin.H{"status": 500, "err": "isNotExist"})
		return
	}

	timeInfo := time.Unix(int64(timeNumber), 0)

	if len(data[fmt.Sprint(timeInfo.Year())]) == 0 {
		data[fmt.Sprint(timeInfo.Year())] = make(map[string]map[string][]DataType)
		data[fmt.Sprint(timeInfo.Year())][timeInfo.Format("01")] = make(map[string][]DataType)
		data[fmt.Sprint(timeInfo.Year())][timeInfo.Format("01")][timeInfo.Format("02")] = []DataType{}
	}

	data[fmt.Sprint(timeInfo.Year())][timeInfo.Format("01")][timeInfo.Format("02")] = append(
		data[fmt.Sprint(timeInfo.Year())][timeInfo.Format("01")][timeInfo.Format("02")],
		DataType{
			Time:    int(timeNumber),
			User:    user,
			Change:  changeNumber,
			Comment: comment,
		},
	)

	updataData(name, data)
	c.JSON(200, gin.H{"status": 200})
}

func getRecord(c *gin.Context) {
	name := c.PostForm("name")
	user := c.PostForm("user")

	var am int = 0
	var in int = 0
	var out int = 0

	data, err := readData(name)
	if err != nil {
		c.JSON(200, gin.H{"status": 500, "err": "isNotExist"})
		return
	}
	m := data[fmt.Sprint(time.Now().Year())][time.Now().Format("01")]
	for _, j := range m {
		for _, jj := range j {
			if user == jj.User || user == "" {
				am += jj.Change
				if jj.Change > 0 {
					in += jj.Change
				} else {
					out -= jj.Change
				}
			}
		}
	}

	var lm int = 0
	lastTime := time.Now().AddDate(0, -1, 0)
	lastM := data[fmt.Sprint(lastTime.Year())][lastTime.Format("01")]
	for _, j := range lastM {
		for _, jj := range j {
			if user == jj.User || user == "" {
				lm += jj.Change
			}
		}
	}

	c.JSON(200, gin.H{"status": 200, "data": m, "am": am, "in": in, "out": out, "lm": lm})
}

func getRecordYear(c *gin.Context) {
	name := c.PostForm("name")
	user := c.PostForm("user")

	var am int = 0
	var in int = 0
	var out int = 0

	data, err := readData(name)
	if err != nil {
		c.JSON(200, gin.H{"status": 500, "err": "isNotExist"})
		return
	}
	m := data[fmt.Sprint(time.Now().Year())]
	for _, j := range m {
		for _, jj := range j {
			for _, jjj := range jj {
				if user == jjj.User || user == "" {
					am += jjj.Change
					if jjj.Change > 0 {
						in += jjj.Change
					} else {
						out -= jjj.Change
					}
				}
			}
		}
	}

	c.JSON(200, gin.H{"status": 200, "data": m, "am": am, "in": in, "out": out})
}
