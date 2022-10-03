// models.login.go

package models

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

type Credential struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordhash"`
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckCredential(username string, password string) bool {
	user := GetUser("SELECT * FROM accounts WHERE username = $1", []any{username})

	bytePassword := []byte(password)
	byteHash := []byte(user[0].PasswordHash)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)

	if err != nil {
		fmt.Println("Err", err.Error())
		return false
	}
	return true
}

func GetUser(sqlStatement string, values []any) []Credential {
	results, err := getRows(sqlStatement, values)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	return convertRowsToCredential(results)
}

func convertRowsToCredential(results *sql.Rows) []Credential {
	credentials := []Credential{}
	for results.Next() {
		var prod Credential
		err := results.Scan(&prod.Id, &prod.Username, &prod.PasswordHash)
		if err != nil {
			panic(err.Error())
		}
		credentials = append(credentials, prod)
	}
	return credentials
}
