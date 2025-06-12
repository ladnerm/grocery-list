package api

import (
	"database/sql"
	"github.com/ladnerm/grocery-list/storage"
	"github.com/ladnerm/grocery-list/templates"
	"github.com/ladnerm/grocery-list/types"
	_ "github.com/lib/pq"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func handlerGetIndex() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
		templates.Temp.Execute(c.Writer, nil)
	}
}

func handlerGetItems() func(*gin.Context) {
	return func(c *gin.Context) {
		var itemArr []types.Item

		query := "SELECT * FROM items;"

		rows, err := db.Query(query)
		if err != nil {
			log.Fatalf("Could not perform query: %v\n", err)
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name string
			var user string
			var location string

			err = rows.Scan(&id, &name, &user, &location)
			if err != nil {
				log.Fatalf("could not scan the row: %v", err)
			}
			itemArr = append(itemArr, *types.NewItem(id, name, user, location))
		}

		c.JSON(http.StatusOK, itemArr)
	}
}

func handlerPostForm() func(*gin.Context) {
	return func(c *gin.Context) {
		newName := c.PostForm("item")
		newUser := c.PostForm("user")
		newLoc := c.PostForm("location")

		var (
			maxid    int
			name     string
			user     string
			location string
		)

		query := "SELECT * FROM items WHERE id = (SELECT MAX(id) from items);"

		row := db.QueryRow(query)
		err := row.Scan(&maxid, &name, &user, &location)
		if err == sql.ErrNoRows {
			maxid = -1
		}

		query = "INSERT INTO items VALUES ($1, $2, $3, $4);"
		_, err = db.Exec(query, maxid+1, newName, newUser, newLoc)
		if err != nil {
			log.Fatal(err)
		}

		c.Redirect(http.StatusSeeOther, "/")
	}
}

func handlerDeleteItem() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		query := "DELETE FROM items WHERE id=$1"
		_, err := db.Exec(query, id)
		if err != nil {
			log.Fatalf("could not delete item with id: %v. ERROR: ", err)
		}
	}
}

func addHandlers(e *gin.Engine) {
	e.GET("/", handlerGetIndex())
	e.GET("/items", handlerGetItems())
	e.POST("/form", handlerPostForm())
	e.DELETE("/delete/:id", handlerDeleteItem())
}

func StartServer() {
	db = storage.InitDB()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	addHandlers(router)
	router.Run() //:8080
}
