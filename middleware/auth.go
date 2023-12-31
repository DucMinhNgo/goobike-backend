package middleware

import (
	"fmt"
	"goobike-backend/common"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		words := strings.Fields(tokenString)
		fmt.Println(words[0])

		if words[0] != "bearer" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			// context.Abort()
			return
		}

		if words[1] == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := common.ValidateToken(words[1])
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
