package handlers

import (
    "net/http"
    "grocery-list/types"
    "fmt"
    "os"
    "encoding/json"

	"github.com/gin-gonic/gin"
)

func HandlerGetIndex() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

func HandlerPostForm() func(*gin.Context) {
	return func(c *gin.Context) {
		str := c.PostForm("item")
		usr := c.PostForm("user")
		loc := c.PostForm("location")

		item := types.NewItem(str, usr, loc)

		fmt.Println(item.Name)
		fmt.Println(item.User)
		fmt.Println(item.Location)

		db, dbErr := os.OpenFile("db/db.json", os.O_RDWR, 0644)

		if dbErr != nil {
			fmt.Printf("DB ERROR! %v", dbErr)
		}

        dcdr := json.NewDecoder(db)
        var itemArr []types.Item
        err := dcdr.Decode(&itemArr)
        if err != nil {
            //do something    
        }
        itemArr = append(itemArr, *item)
		db.Close()

		db, dbErr = os.Create("db/db.json")
		defer db.Close()
		if dbErr != nil {
			fmt.Printf("DB ERROR! %v", dbErr)
		}
		encdr := json.NewEncoder(db)

		if err := encdr.Encode(itemArr); err != nil {
			fmt.Print(err)
		}

		c.Redirect(http.StatusSeeOther, "/")
	}
}

