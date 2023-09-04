package main

import (
    "fmt"
    "net/http"
    "time"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "github.com/joho/godotenv"
    "os"
)

type TodoItem struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    Status string `json:"status"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type TodoItemCreation struct {
    Title string `json:"title" gorm:"column:title"`
    Description string `json:"description" gorm:"column:description"`
    Status string `json:"status" gorm:"column:status"`
}

func (TodoItemCreation) TableName() string { return "todo_items" }

func testTodoItem() {
    // UTC: lấy thời gian gốc (múi giờ)
    now := time.Now().UTC()

    item := TodoItem{
        Id:          1,
        Title:       "Test Title",
        Description: "Test Description",
        Status:      "Doing",
        CreatedAt:   &now,
        UpdatedAt:   &now,
    }

    jsonData, err := json.Marshal(item)

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(jsonData))

    const jsonStr = "{\"id\":1,\"title\":\"Test Title\",\"description\":\"Test Description\",\"status\":\"Doing\",\"created_at\":\"2023-09-04T01:15:16.199774Z\",\"updated_at\":\"2023-09-04T01:15:16.199774Z\"}"

    var item2 TodoItem

    json.Unmarshal([]byte(jsonStr), &item2)

    fmt.Println(item2)
}

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
    dsn := os.Getenv("MYSQL_DATABASE_CONNECTION")

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if (err != nil) {
        log.Fatal(err)
    }

    fmt.Println(db)

    now := time.Now().UTC()

    item := TodoItem{
        Id:          1,
        Title:       "Test Title",
        Description: "Test Description",
        Status:      "Doing",
        CreatedAt:   &now,
        UpdatedAt:   &now,
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
            items.GET("")
            items.GET("/:id")
            items.PUT("/:id")
            items.DELETE("/:id")
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
        if err := c.ShouldBind(&data); err != nil {
           c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })

            return;
        }

        if err := db.Create(&data).Error; err != nil {
            // internal server error
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })

            return; 
        }

        c.JSON(http.StatusOK, gin.H{
            "data": data,
        }) 
    }
} 