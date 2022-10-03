// handlers.login.go

package handlers

import (
	"net/http"

	"github.com/Sevitama/gin-poc/models"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	authenticated := models.CheckCredential(username, password)

	c.HTML(http.StatusOK, "login.html", gin.H{"authenticated": authenticated})
}
