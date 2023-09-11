package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Enum
type ItemStatus int

const (
	ItamStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemstatuses = [3]string{"doing", "done", "deleted"}

func (item *ItemStatus) String() string {
	return allItemstatuses[*item]
}

func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range allItemstatuses {
		if allItemstatuses[i] == s {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), errors.New("invalid status")
}

// Doc du lieu mang len Itemstatus
func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	v, err := parseStr2ItemStatus(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	*item = v

	return nil
}

// Itemstatus veef duwx lieeuj
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}

// Ho tro Itemstatus thanh JsonValue
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

// JsonValue sang Items status
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	itemValue, err := parseStr2ItemStatus(str)
	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      ItemStatus `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItemCreation) TableName() string { return "todo_items" }

// * string (khi truyền ” vẫn cập nhật)
type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItemUpdate) TableName() string { return "todo_items" }

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

	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "Test Title",
		Description: "Test Description",
		// Status:      "Doing",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

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
			items.POST("", CreateItem(db))
			items.GET("", GetList(db))
			items.GET("/:id", GetItem(db))
			items.PUT("/:id", UpdateItem(db))
			items.DELETE("/:id", DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})

	r.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreation
		// UnmarshalJSON func
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// Value func
		if err := db.Create(&data).Error; err != nil {
			// internal server error
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem

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

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func GetList(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var result []TodoItem

		if err := db.Find(&result).Error; err != nil {
			if err != nil {
				// internal server error
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemUpdate

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

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
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
		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
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

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
