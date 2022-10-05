package handlers

import (
  "net/http"

  "github.com/Sevitama/gin-poc/models"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
)

type SignInInput struct {
  Username string `json:"username" binding:"required"`
  Password string `json:"password" binding:"required"`
}

func SignIn(c *gin.Context) {
  var input SignInInput

  if err := c.ShouldBindWith(&input, binding.Form); err != nil {
    c.HTML(http.StatusBadRequest,"error.html", gin.H{"error": err.Error()})
    return
  }

  token, err := models.LoginCheck(input.Username, input.Password)

  if err != nil {
    c.HTML(http.StatusBadRequest,"error.html", gin.H{"error": "username or password is incorrect."})
    return
  }

  c.HTML(http.StatusOK, "login.html", gin.H{"authenticated": "authenticated", "token": token})
}
