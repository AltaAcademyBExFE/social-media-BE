package domain

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Core struct {
	gorm.Model
	Body   string
	PostID int
	UserID int
}

type Cores struct {
	gorm.Model
	Body string
	Name string
}

type Repository interface {
	Insert(newComment Core) (Cores, error)
	Del(ID int) error
}

type Service interface {
	Create(newComment Core) (Cores, error)
	Delete(ID int) error
}

type Handler interface {
	CreateComment() echo.HandlerFunc
	DeleteComment(ID int) echo.HandlerFunc
}
