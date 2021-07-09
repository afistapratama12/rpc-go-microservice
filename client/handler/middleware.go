package handler

import (
	"microGoLearn/client/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware(auth auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("Authorization")

		if len(key) == 0 || key == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		token, err := auth.ValidateToken(key)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		userID := claim["user_id"].(string)
		userEmail := claim["email"].(string)

		c.Set("currentUserId", userID)
		c.Set("currentUserEmail", userEmail)
	}
}
