package handlers

import (
  "net/http"

  "github.com/Sevitama/gin-poc/utils/token"
  "github.com/Sevitama/gin-poc/models"
  "github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
  user_id, err := token.ExtractTokenID(c)

  if err != nil {
    c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
    return
  }

  credential := models.GetSingleCredential("SELECT * FROM accounts WHERE id = $1", []any{user_id})

  c.HTML(http.StatusOK, "data.html", gin.H{"Id": user_id, "username": credential.Username, "passwordhash": credential.PasswordHash})
}
