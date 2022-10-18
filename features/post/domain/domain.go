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

type Cores struct {
	gorm.Model
	Body   string
	Images string
	Name   string
}

type Repository interface {
	Show() ([]Cores, error)
	My(ID int) ([]Cores, error)
	Spesific(ID int) ([]Cores, error)
	Insert(newPost Core) (Cores, error)
	Update(ID int, updatePost Core) (Cores, error)
	Del(ID int) error
}

type Service interface {
	ShowAll() ([]Cores, error)
	ShowMy(ID int) ([]Cores, error)
	ShowSpesific(ID int) ([]Cores, error)
	Create(newPost Core) (Cores, error)
	Edit(ID int, updatePost Core) (Cores, error)
	Delete(ID int) error
}

type Handler interface {
	ShowAllPost() echo.HandlerFunc
	ShowMyPost(ID int) echo.HandlerFunc
	ShowSpesificPost(ID int) echo.HandlerFunc
	CreatePost() echo.HandlerFunc
	EditPost(ID int) echo.HandlerFunc
	DeletePost(ID int) echo.HandlerFunc
}
