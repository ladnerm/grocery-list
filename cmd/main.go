package main

import (
	"grocery-list/handlers"

	"github.com/gin-gonic/gin"
)

func addHandlers(e *gin.Engine) {
	x := handlers.HandlerGetIndex()
	y := handlers.HandlerPostForm()
	z := handlers.HandlerGetItems()
	// URL, WHAT TO DO WITH THE INFORMATION
	e.GET("/", x)
	e.GET("/items", z)
	e.POST("/form", y)
}

// TODO:
// total cost of all products = sum of each estimated price
func main() {
	router := gin.Default()

	// Load HTML Templates
	router.LoadHTMLGlob("templates/*")

	addHandlers(router)

	router.Run() //:8080
}
