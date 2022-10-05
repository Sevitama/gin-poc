// main.go

package main

import (
	"github.com/Sevitama/gin-poc/handlers"
	"github.com/Sevitama/gin-poc/middlewares"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/secureSQLi/", handlers.GetArticlesSecureSQLi)
	router.GET("/article/insecureSQLi/", handlers.GetArticlesInsecureSQLi)
	router.POST("/signin",handlers.SignIn)

	router.Use(middlewares.JwtAuthMiddleware())
	router.GET("/data", handlers.GetData)
	// Start serving the application
	router.Run()

}
