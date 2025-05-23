package main

import (
    "grocery-list/handlers"

	"github.com/gin-gonic/gin"
)

func addHandlers(e *gin.Engine) {
    x := handlers.HandlerGetIndex()
	y := handlers.HandlerPostForm()
	// URL, WHAT TO DO WITH THE INFORMATION
	e.GET("/", x)
	e.POST("/form", y)
}

func main() {
	router := gin.Default()

	// Load HTML Templates
	router.LoadHTMLGlob("templates/*")

    addHandlers(router)

	router.Run() //:8080
}
