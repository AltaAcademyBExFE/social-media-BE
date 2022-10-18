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

type Repository interface {
	Insert(newComment Core) (Core, error)
	Del(ID int) error
}

type Service interface {
	Create(newComment Core) (Core, error)
	Delete(ID int) error
}

type Handler interface {
	CreateComment() echo.HandlerFunc
	DeleteComment(ID int) echo.HandlerFunc
}
