package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Location string `json:"location"`
}

func newItem(name string, user string, location string) *Item {
	return &Item{
		Name:     name,
		User:     user,
		Location: location,
	}
}

func indexGet() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

func formGet() func(*gin.Context) {
	return func(c *gin.Context) {
		str := c.PostForm("item")
		usr := c.PostForm("user")
		loc := c.PostForm("location")

		item := newItem(str, usr, loc)

		fmt.Println(item.Name)
		fmt.Println(item.User)
		fmt.Println(item.Location)

		c.Redirect(http.StatusSeeOther, "/")
	}
}

func main() {
	router := gin.Default()

	// Load HTML Templates
	router.LoadHTMLGlob("templates/*")

	var x func(*gin.Context) = indexGet()
	y := formGet()
	// URL, WHAT TO DO WITH THE INFORMATION
	router.GET("/", x)
	router.POST("/form", y)
	router.Run() //:8080
}
