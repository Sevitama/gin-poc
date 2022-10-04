// handlers.article.go

package handlers

import (
	"net/http"

	"github.com/utrack/gin-csrf"
	"github.com/Sevitama/gin-poc/models"
	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context, getArticles func(string) []models.Article) {
	title := c.Query("title")
	articles := getArticles(title)
	c.HTML(http.StatusOK, "article.html", gin.H{"payload": articles})
}

func GetArticlesSecureSQLi(c *gin.Context) {
	GetArticles(c, models.GetArticlesByTitleSecure)
}

func GetArticlesInsecureSQLi(c *gin.Context) {
	GetArticles(c, models.GetArticlesByTitleInsecure)
}

func GetToken(c *gin.Context) {
	c.HTML(http.StatusOK, "csrf.html", gin.H{"token": csrf.GetToken(c)})
}

func PostToken(c *gin.Context) {
	c.HTML(http.StatusOK, "csrf.html", gin.H{"token": csrf.GetToken(c)})
}
