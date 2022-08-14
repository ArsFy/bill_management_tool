package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// HTTP
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")
		c.Next()
	}
}

func status(c *gin.Context) {
	c.JSON(200, gin.H{"status": 200, "version": config["version"].(string)})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(cors())
	router.Use(gin.Recovery())
	// 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": 404})
	})
	// Status
	router.POST("/api/status", status)
	router.GET("/api/status", status)

	// API
	router.POST("/api/obj_list", objList)
	router.POST("/api/create_obj", createObj)

	// Run
	fmt.Printf("BMTool v%s Starting :%s ...\n", config["version"].(string), config["port"].(string))
	router.Run(":" + config["port"].(string))
}