package api

import (
	"github.com/ladnerm/grocery-list/templates"
	"github.com/ladnerm/grocery-list/types"
	"github.com/ladnerm/grocery-list/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	addHandlers(router)
	router.Run() //:8080
}

func addHandlers(e *gin.Engine) {
	e.GET("/", handlerGetIndex())
	e.GET("/items", handlerGetItems())
	e.POST("/form", handlerPostForm())
	e.DELETE("/:id", handlerDeleteItem())
}

func handlerGetIndex() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
		templates.Temp.Execute(c.Writer, nil)
	}
}

func handlerGetItems() func(*gin.Context) {
	return func(c *gin.Context) {
		var itemArr []types.Item
		err := util.ItemsFromDB(&itemArr)
		if err != nil {
			log.Fatalf("Could not open db: %v\n", err)
		}

		c.JSON(http.StatusOK, itemArr)
	}
}

func handlerPostForm() func(*gin.Context) {
	return func(c *gin.Context) {
		str := c.PostForm("item")
		usr := c.PostForm("user")
		loc := c.PostForm("location")
		item := types.NewItem(str, usr, loc)

		var itemArr []types.Item

		err := util.ItemsFromDB(&itemArr)
		if err != nil {
			log.Fatalf("FORM: could not get items from DB: %v\n", err)
		}

		if len(itemArr) != 0 {
			item.ID = itemArr[len(itemArr)-1].ID + 1
		}
		itemArr = append(itemArr, *item)

		err = util.WriteToDB(itemArr)
		if err != nil {
			log.Fatalf("FORM: could not write to DB: %v\n", err)
		}

		c.Redirect(http.StatusSeeOther, "/")
	}
}

func handlerDeleteItem() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var itemArr []types.Item

		err := util.ItemsFromDB(&itemArr)
		if err != nil {
			log.Fatalf("FORM: could not get items from DB: %v\n", err)
		}

		var newItemArr []types.Item

		for _, v := range itemArr {
			idStr, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal("could not convert id to string")
			}
			if v.ID != idStr {
				newItemArr = append(newItemArr, v)
			}
		}

		util.WriteToDB(newItemArr)
	}
}
