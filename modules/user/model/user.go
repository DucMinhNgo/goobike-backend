package model

import (
	"errors"
	"goobike-backend/common"

	"golang.org/x/crypto/bcrypt"
)

const (
	EntityName = "User"
)

var (
	ErrEmailIsBlank    = errors.New("enail cannot be blank")
	ErrPasswordIsBlank = errors.New("password cannot be blank")
	ErrNameIsBlank     = errors.New("name cannot be blank")
	ErrUsernameIsBlank = errors.New("username cannot be blank")
)

type User struct {
	common.SQLModel
	Name     string `json:"name" gorm:"column:name"`
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (User) TableName() string { return "users" }

type UserCreation struct {
	common.SQLModel
	Name     string `json:"name" gorm:"column:name"`
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (user *UserCreation) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (UserCreation) TableName() string { return "users" }
