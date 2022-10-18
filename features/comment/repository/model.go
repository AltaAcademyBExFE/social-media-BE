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

func FromDomain(do domain.Core) Comment {
	return Comment{
		Model:  gorm.Model{ID: do.ID},
		Body:   do.Body,
		PostID: do.PostID,
		UserID: do.UserID,
	}
}

func ToDomain(c Comment) domain.Core {
	return domain.Core{
		Model:  gorm.Model{ID: c.ID},
		Body:   c.Body,
		PostID: c.PostID,
		UserID: c.UserID,
	}
}
