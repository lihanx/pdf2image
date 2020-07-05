package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	InitWorkdir()

	router := gin.Default()
	router.StaticFS("/data", http.Dir("DATA"))
	router.StaticFS("/statics", http.Dir("statics"))
	router.LoadHTMLGlob("./templates/*")
	router.GET("/index", Index)
	router.POST("/PDF2ZIP", PDF2Zip)
	router.POST("/PDF2JPG", PDF2JPG)

	router.Run(":10086")
}