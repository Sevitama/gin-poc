// main.go

package main

import (
	"net/http"
	"github.com/Sevitama/gin-poc/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/utrack/gin-csrf"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
      c.HTML(http.StatusBadRequest, "csrf.html", gin.H{"token": "Yikes: " + csrf.GetToken(c)})
			c.Abort()
		},
	}))

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/secureSQLi/", handlers.GetArticlesSecureSQLi)
	router.GET("/article/insecureSQLi/", handlers.GetArticlesInsecureSQLi)
  router.GET("/token/", handlers.GetToken)
  router.POST("/token/", handlers.PostToken)
	router.POST("/signin", handlers.SignIn)
	// Start serving the application
	router.Run()

}
