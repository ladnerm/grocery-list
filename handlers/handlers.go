package handlers

import (
	"encoding/json"
	"fmt"
	"grocery-list/templates"
	"grocery-list/types"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandlerGetIndex() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
		templates.Temp.Execute(c.Writer, nil)
	}
}

func HandlerGetItems() func(*gin.Context) {
	return func(c *gin.Context) {
		db, dbErr := os.OpenFile("db/db.json", os.O_RDONLY, 0644)
		if dbErr != nil {
			fmt.Printf("DB ERROR! %v", dbErr)
		}
		dcdr := json.NewDecoder(db)

		var itemArr []types.Item
		err := dcdr.Decode(&itemArr)
		if err != nil {
			//do something
		}
		c.JSON(http.StatusOK, itemArr)
	}
}

func HandlerPostForm() func(*gin.Context) {
	return func(c *gin.Context) {
		str := c.PostForm("item")
		usr := c.PostForm("user")
		loc := c.PostForm("location")

		item := types.NewItem(str, usr, loc)

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

		if len(itemArr) != 0 {
			item.ID = itemArr[len(itemArr)-1].ID + 1
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

func HandlerDeleteItem() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

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
		db.Close()

		var newItemArr []types.Item

		for _, v := range itemArr {
			iid, err := strconv.Atoi(id)
			if err != nil {
				//Do something
			}
			if v.ID != iid {
				newItemArr = append(newItemArr, v)
			}
		}

		db, dbErr = os.Create("db/db.json")
		defer db.Close()
		if dbErr != nil {
			fmt.Printf("DB ERROR! %v", dbErr)
		}
		encdr := json.NewEncoder(db)
		err = encdr.Encode(newItemArr)
		if err != nil {
			fmt.Print(err)
		}

	}
}
