package middleware

import (
	"goobike-backend/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery() func(*gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					c.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrorInternal(err))
				}
				panic(r)
			}

		}()
		c.Next()
	}
}
