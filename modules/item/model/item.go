package model

import (
	"errors"
	"goobike-backend/common"
)

var (
	ErrTitleIsBlank = errors.New("title cannot be blank")
)

type TodoItem struct {
	// embeding Struct (khong phai ke thua)
	common.SQLModel
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      ItemStatus `json:"status" gorm:"column:status"`
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
