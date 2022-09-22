package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const dbuser = "postgres"
const dbpass = "postgres"
const dbname = "newsplatform"
const dbip = "0.0.0.0"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetAllArticles() []Article {
	connStr := "postgres://" + dbuser + ":" + dbpass + "@" + dbip + "/" + dbname + "?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	println("here")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM article")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	articles := []Article{}

	for results.Next() {
		var prod Article
		err = results.Scan(&prod.ID, &prod.Title, &prod.Content)
		if err != nil {
			panic(err.Error())
		}
		articles = append(articles, prod)
	}
	return articles
}

func GetArticleByID(id int) *Article {
	connStr := "postgres://" + dbuser + ":" + dbpass + "@" + dbip + "/" + dbname + "?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	prod := &Article{}

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM article where id = $1", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&prod.ID, &prod.Title, &prod.Content)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}
	return prod
}