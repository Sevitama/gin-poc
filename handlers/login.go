// handlers.login.go

package handlers

import (
	"fmt"
	"net/http"

	"github.com/Sevitama/gin-poc/models"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("DATA:")
	fmt.Println(password)
	fmt.Println(username)
	authenticated := models.CheckCredential(username, password)

	c.HTML(http.StatusOK, "login.html", gin.H{"authenticated": authenticated})
}
