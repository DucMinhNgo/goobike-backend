package ginuser

import (
	"goobike-backend/common"
	"goobike-backend/modules/user/biz"
	"goobike-backend/modules/user/model"
	"goobike-backend/modules/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GenerateToken(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var request model.TokenRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewGetUserBiz(store)

		user, err := business.GetUserByEmail(c.Request.Context(), request.Email)

		if err != nil {
			// internal server error
			c.JSON(http.StatusBadRequest, err)

			return
		}

		credentialError := user.CheckPassword(request.Password)

		if credentialError != nil {
			c.JSON(http.StatusUnauthorized, common.ErrInvalidPassword(credentialError))
			c.Abort()
			return
		}
		tokenString, err := common.GenerateJWT(user.Email, user.Username)
		if err != nil {
			// internal server error
			c.JSON(http.StatusBadRequest, err)
			return
		}

		data := model.TokenResponse{Token: tokenString}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
