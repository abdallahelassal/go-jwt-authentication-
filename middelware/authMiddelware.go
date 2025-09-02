package middelware

import (
	"net/http"
	"strings"
	"github.com/abdallahelassal/go-jwt-authentication-.git/helpers"
	"github.com/gin-gonic/gin"
)

func AuthMiddelware()gin.HandlerFunc{
	return  func(c *gin.Context) {
		var jwtToken string

		if authHeader := c.GetHeader("Authorization"); strings.HasPrefix(jwtToken,"Bearer "){
			jwtToken = strings.TrimPrefix(authHeader,"Bearer ")
		}else {
			cookie , err := c.Cookie("jwt")
			if err != nil {
				jwtToken = cookie
			}
		}
		if jwtToken == ""{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token "})
			return
		}
		climes , err := helpers.ValidateToken(jwtToken , "access")
		if err != nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token "})
			return 
		}
		c.Set("userID", climes.UserID)
		c.Set("email", climes.Email)
		c.Next()
	}
}
