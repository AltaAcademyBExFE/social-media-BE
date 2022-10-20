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

type Comes struct {
	gorm.Model
	Body   string
	Name   string
	PostID int
}

type Repository interface {
	Show() ([]Cores, []Comes, error)
	My(ID int) ([]Cores, []Comes, error)
	Spesific(ID int) ([]Cores, []Comes, error)
	Insert(newPost Core) (Cores, error)
	Update(ID int, updatePost Core) (Cores, error)
	Del(ID int) error
}

type Service interface {
	ShowAll() ([]Cores, []Comes, error)
	ShowMy(ID int) ([]Cores, []Comes, error)
	ShowSpesific(ID int) ([]Cores, []Comes, error)
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
