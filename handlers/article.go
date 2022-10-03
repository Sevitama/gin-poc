// handlers.article.go

package handlers

import (
	"net/http"

	"github.com/Sevitama/gin-poc/models"
	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context, getArticles func(string) []models.Article) {
	title := c.Query("title")
	articles := getArticles(title)
	c.HTML(http.StatusOK, "article.html", gin.H{"payload": articles})
}

func GetArticlesSecureSQLInjection(c *gin.Context) {
	GetArticles(c, models.GetArticlesByTitleSecureSQLInjection)
}

func GetArticlesInsecureSQLInjection(c *gin.Context) {
	GetArticles(c, models.GetArticlesByTitleInsecureSQLInjection)
}
