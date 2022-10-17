package domain

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Core struct {
	gorm.Model
	Body   string
	Images string
	UserID int
}

type Repository interface {
	Show() ([]Core, error)
	My(ID int) ([]Core, error)
	Insert(newPost Core) (Core, error)
	Update(updatePost Core) (Core, error)
	Del(ID int) error
}

type Service interface {
	ShowAll() ([]Core, error)
	ShowMy(ID int) ([]Core, error)
	Create(newPost Core) (Core, error)
	Edit(updatePost Core) (Core, error)
	Delete(ID int) error
}

type Handler interface {
	ShowAllPost() echo.HandlerFunc
	ShowMyPost(ID int) echo.HandlerFunc
	CreatePost() echo.HandlerFunc
	EditPost() echo.HandlerFunc
	DeletePost(ID int) echo.HandlerFunc
}
