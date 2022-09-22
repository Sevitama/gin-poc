// handlers.article.go

package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sevitama/gin-poc/models"
	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	globalArticles := models.GetAllArticles()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": globalArticles,
		},
	)

}

func GetArticle(c *gin.Context) {
	// Check if the article ID is valid
	fmt.Println(strconv.Atoi(c.Param("article_id")))

	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article := models.GetArticleByID(articleID); err == nil {

			// Call the HTML method of the Context to render a template
			c.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"article.html",
				// Pass the data that the page uses
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
