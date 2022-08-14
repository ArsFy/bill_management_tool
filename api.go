package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

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
		WriteFile(path.Join(wd, "./data/", name+".json"), "[]")
		c.JSON(200, gin.H{"status": 200, "info": "ok"})
	}
}
