package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func convertRowsToArticles(results *sql.Rows) []Article {
	articles := []Article{}

	for results.Next() {
		var prod Article
		err := results.Scan(&prod.ID, &prod.Title, &prod.Content)
		if err != nil {
			panic(err.Error())
		}
		articles = append(articles, prod)
	}
	return articles
}

func GetArticles(sqlStatement string, values []any) []Article {
	results, err := getRows(sqlStatement, values)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	return convertRowsToArticles(results)
}

func GetArticlesByTitleInsecureSQLi(title string) []Article {
	// db.Query(fmt.Sprintf("SELECT * FROM article WHERE title = '%s'", title))
	return GetArticles(fmt.Sprintf("SELECT * FROM article WHERE title = '%s'", title), []any{})
}

func GetArticlesByTitleSecureSQLi(title string) []Article {
	// db.Query("SELECT * FROM article WHERE title = $1", title)
	return GetArticles("SELECT * FROM article WHERE title = $1", []any{title})
}

func GetArticlesByTitleSecureXSS(title string) []Article {
	return GetArticles("SELECT * FROM article WHERE title = $1", []any{title})
}
