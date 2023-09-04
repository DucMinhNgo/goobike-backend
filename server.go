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

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
        "message": item,
        })
    })

    r.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}