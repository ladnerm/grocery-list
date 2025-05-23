package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
    "github.com/ladnerm/grocery-list/types"

	"github.com/gin-gonic/gin"
)


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

		item := types.NewItem(str, usr, loc)

		fmt.Println(item.Name)
		fmt.Println(item.User)
		fmt.Println(item.Location)

		db, dbErr := os.OpenFile("db/db.json", os.O_APPEND|os.O_WRONLY, 0644)
		defer db.Close()
		if dbErr != nil {
			fmt.Printf("DB ERROR! %v", dbErr)
		}
		encdr := json.NewEncoder(db)
		if err := encdr.Encode(*item); err != nil {
			fmt.Print(err)
			fmt.Print("here")
		}

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
