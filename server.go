package main

import (
	"fmt"
	"goobike-backend/common"
	"goobike-backend/modules/item/model"
	ginitem "goobike-backend/modules/item/transport/gin"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func testTodoItem() {
// 	// UTC: lấy thời gian gốc (múi giờ)
// 	now := time.Now().UTC()

// 	item := TodoItem{
// 		Id:          1,
// 		Title:       "Test Title",
// 		Description: "Test Description",
// 		// Status:      "Doing",
// 		CreatedAt: &now,
// 		UpdatedAt: &now,
// 	}

// 	jsonData, err := json.Marshal(item)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(jsonData))

// 	const jsonStr = "{\"id\":1,\"title\":\"Test Title\",\"description\":\"Test Description\",\"status\":\"Doing\",\"created_at\":\"2023-09-04T01:15:16.199774Z\",\"updated_at\":\"2023-09-04T01:15:16.199774Z\"}"

// 	var item2 TodoItem

// 	json.Unmarshal([]byte(jsonStr), &item2)

// 	fmt.Println(item2)
// }

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("MYSQL_DATABASE_CONNECTION")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	// now := time.Now().UTC()

	// item := model.TodoItem{
	// 	Id:          1,
	// 	Title:       "Test Title",
	// 	Description: "Test Description",
	// 	// Status:      "Doing",
	// 	CreatedAt: &now,
	// 	UpdatedAt: &now,
	// }

	r := gin.Default()

	// CRUD
	// POST /v1/items (create new item)
	// GET /v1/items?page=1 (list item)
	// GET /v1/items/:id (get item detail by id)
	// (PUT | PATCH) /v1/items/:id (update item by id)
	// DELETE /v1/items/:id (delete item by id)
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", GetList(db))
			items.GET("/:id", GetItem(db))
			items.PUT("/:id", UpdateItem(db))
			items.DELETE("/:id", DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItem

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			// internal server error
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		data.Id = id

		if err := db.First(&data).Error; err != nil {
			if err != nil {
				// internal server error
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

				return
			}
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func GetList(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Process()

		var result []model.TodoItem

		db = db.Where("status <> ?", "deleted")

		if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Order("id desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Find(&result).Error; err != nil {
			if err != nil {
				// internal server error
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

				return
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			// internal server error
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			if err != nil {
				// internal server error
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

				return
			}
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			// internal server error
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		// hard delete
		// if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		// soft delete
		if err := db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
			"status": "Deleted",
		}).Error; err != nil {
			if err != nil {
				// internal server error
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

				return
			}
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
