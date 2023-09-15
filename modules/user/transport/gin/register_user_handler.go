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

func RegisterUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var user model.UserCreation
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := user.HashPassword(user.Password); err != nil {
			c.JSON(http.StatusInternalServerError, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewRegisterUserBiz(store)

		if err := business.RegisterNewUser(c.Request.Context(), &user); err != nil {
			// 	// internal server error
			c.JSON(http.StatusBadRequest, err)

			return
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(user))
	}
}
