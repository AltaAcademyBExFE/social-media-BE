package repository

import (
	"sosmed/features/comment/domain"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Body   string
	PostID int
	UserID int
}

type CommentIt struct {
	gorm.Model
	Body string
	Name string
}

func FromDomain(do domain.Core) Comment {
	return Comment{
		Model:  gorm.Model{ID: do.ID},
		Body:   do.Body,
		PostID: do.PostID,
		UserID: do.UserID,
	}
}

func ToDomain(c CommentIt) domain.Cores {
	return domain.Cores{
		Model: gorm.Model{ID: c.ID, CreatedAt: c.CreatedAt},
		Body:  c.Body,
		Name:  c.Name,
	}
}
