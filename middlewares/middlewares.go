package middlewares

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/Sevitama/gin-poc/utils/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {

  return func(c *gin.Context) {
    err := token.TokenValid(c)
    if err != nil {
      c.HTML(http.StatusUnauthorized,"error.html", gin.H{"error": "Unauthorized"})
      c.Abort()
      return
    }
    c.Next()
  }
}
