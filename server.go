package main

import "github.com/gin-gonic/gin"
import (
    "url-shortener/test/test"
)

func main() {
	r := gin.Default()
	r.GET("/ping", testGET)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}