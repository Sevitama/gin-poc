// main.go

package main

import (
	"github.com/Sevitama/gin-poc/handlers"
	"github.com/Sevitama/gin-poc/models"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var globalArticles []models.Article

func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetArticle)
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	// Start serving the application
	router.Run(":8080")

}
