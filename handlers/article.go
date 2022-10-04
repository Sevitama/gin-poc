// handlers.article.go

package handlers

import (
	"net/http"
  "html/template"

	"github.com/Sevitama/gin-poc/models"
	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context, getArticles func(string) []models.Article) {
	title := c.Query("title")
	articles := getArticles(title)
	c.HTML(http.StatusOK, "article.html", gin.H{"payload": articles})
}

func GetArticlesSecureSQLi(c *gin.Context) {
	GetArticles(c, models.GetArticlesByTitleSecureSQLi)
}

func GetArticlesInsecureSQLi(c *gin.Context) {
	GetArticles(c, models.GetArticlesByTitleInsecureSQLi)
}

func GetArticlesSecureXSS(c *gin.Context) {
  value := c.Query("title")
	c.HTML(http.StatusOK, "xss_secure.html", gin.H{"value": value})
}

func GetArticlesInsecureXSS(c *gin.Context) {
  value := c.Query("title")
	c.HTML(http.StatusOK, "xss_insecure.html", gin.H{"value": template.JS(value)})
}
