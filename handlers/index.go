// handlers.article.go

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
