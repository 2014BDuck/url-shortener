package test

import "github.com/gin-gonic/gin"

func testGET(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}