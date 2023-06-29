package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type document struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var documents = []document{
	{ID: "1", Title: "First document"},
	{ID: "2", Title: "Second document"},
}

func getDocuments(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, documents)
}

func main() {
	fmt.Println("Gin API")
	router := gin.Default()

	router.GET("/api/v1/documents", getDocuments)

	router.Run("localhost:8888")
}
