package models

import (
  "database/sql"
  "fmt"
)

const dbuser = "postgres"
const dbpass = "postgres"
const dbname = "newsplatform"
const dbip = "0.0.0.0"

func getRows(sqlStatement string, values []any) (*sql.Rows, error) {
  connStr := "postgres://" + dbuser + ":" + dbpass + "@" + dbip + "/" + dbname + "?sslmode=disable"
  db, err := sql.Open("postgres", connStr)

  if err != nil {
    fmt.Println("Err", err.Error())
    return nil, err
  }
  defer db.Close()

  return db.Query(sqlStatement, values...)
}

func GetRecord(sqlStatement string, values []any) (*sql.Row) {
  connStr := "postgres://" + dbuser + ":" + dbpass + "@" + dbip + "/" + dbname + "?sslmode=disable"
  db, err := sql.Open("postgres", connStr)

  if err != nil {
    fmt.Println("Err", err.Error())
    return nil
  }
  defer db.Close()

  row := db.QueryRow(sqlStatement, values...)
  return row
}
