// handlers.article.go

package handlers

import (
  "net/http"

  "github.com/Sevitama/gin-poc/models"
  "github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html", gin.H{})
}

func GetArticles(c *gin.Context, getArticles func(string) []models.Article) {
  title := c.Query("title")
  articles := getArticles(title)
  c.HTML( http.StatusOK, "article.html", gin.H{"payload": articles})
}

func GetArticlesSecure(c *gin.Context) {
  GetArticles(c, models.GetArticlesByTitleSecure)
}

func GetArticlesInsecure(c *gin.Context) {
  GetArticles(c, models.GetArticlesByTitleInsecure)
}
