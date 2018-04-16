package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

import . "app/domain"

var db = initDb()

func getTodos(c *gin.Context) {
	// fetch all rows
	var todos []Todo
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	// fetch one row
	id := c.Params.ByName("id")
	var todo Todo
	result := db.First(&todo, id)

	if result.RecordNotFound() {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func createTodo(c *gin.Context) {
	var todo Todo
	// check for errors
	if c.BindJSON(&todo) == nil {
		t := NewTodo(todo.Title)
		if db.NewRecord(t) {
			db.Create(&t)
			c.JSON(http.StatusCreated, t)
		} else {
			c.JSON(http.StatusForbidden, t)
		}
	} else {
		c.JSON(http.StatusNotAcceptable, "Not acceptable")
	}
}

func updateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	result := db.First(&todo, id)

	if result.RecordNotFound() {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	if c.BindJSON(&todo) == nil {
		db.Update(&todo)
		c.JSON(http.StatusOK, todo)
		return
	}

}

func deleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")

	// delete row by PK
	log.Println(id)

	// fetch one row

	var todo Todo
	result := db.First(&todo, id).Delete(&todo)

	if result.RecordNotFound() {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func main() {
	//defer db.Close()
	log.Println("Main running")
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.GET("/", getTodos)
			todos.POST("/", createTodo)
			todos.PUT("/:id", updateTodo)
			todos.DELETE("/:id", deleteTodo)
			todos.GET("/:id", getTodo)
		}
	}
	r.Run(":3001")
}

func initDb() *gorm.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		"postgres",
		os.Getenv("POSTGRES_PASSWORD"),
		"postgres",
		"5432",
		"postgres",
	)

	db, err := gorm.Open("postgres", dbinfo)
	checkErr(err, "gorm.Open failed")

	db.AutoMigrate(&Todo{})

	return db
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
