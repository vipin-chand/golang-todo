package main

import (
	// "fmt"
	// "fmt"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
}

var Books = []book{
	{ID : "1",Name : "Some Name",Author : "Some author"},
	{ID : "2",Name : "Some Name",Author : "Some author"},
	{ID : "3",Name : "Some Name",Author : "Some author"},
}


func main(){
	router := gin.Default()
	router.GET("/books", getAllBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)

	router.Run("localhost:8090")
}

func getAllBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, Books)
}

func createBook(c *gin.Context){
	var newBook book
	
	if err := c.BindJSON(&newBook); err != nil{
		return
	}

	Books = append(Books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)
}

func bookById(c *gin.Context){
	id := c.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		return
	}
	
	c.IndentedJSON(http.StatusOK, book)
}

func getBookByID(id string) (*book, error){
 
	for key, book := range Books{
		if id == book.ID {
			return &Books[key], nil
		}
	}

	return nil, errors.New("book not found")
}

