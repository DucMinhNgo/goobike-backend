package ginitem

import (
	"goobike-backend/common"
	"goobike-backend/modules/item/biz"
	"goobike-backend/modules/item/model"
	"goobike-backend/modules/item/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreation
		// UnmarshalJSON func
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest((err)))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)

		// Value func
		if err := business.CreateNewItem(c.Request.Context(), &data); err != nil {
			// 	// internal server error
			c.JSON(http.StatusBadRequest, err)

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
