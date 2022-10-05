// models.login.go

package models

import (
  "golang.org/x/crypto/bcrypt"
  "github.com/Sevitama/gin-poc/utils/token"
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

func VerifyPassword(password,hashedPassword string) error {
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string,error) {

  user := GetSingleCredential("SELECT * FROM accounts WHERE username = $1", []any{username})

  err := VerifyPassword(password, user.PasswordHash)

  if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
   return "", err
  }

  token,err := token.GenerateToken(user.Id)

  if err != nil {
   return "",err
  }

  return token,nil
}

func GetSingleCredential(sqlStatement string, values []any) Credential {
  credential := Credential{}
  result := GetRecord(sqlStatement, values)
  err := result.Scan(&credential.Id, &credential.Username, &credential.PasswordHash)

  if err != nil {
    panic(err.Error())
  }

  return credential
}
