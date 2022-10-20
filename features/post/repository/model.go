package repository

import (
	"sosmed/features/post/domain"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Body   string
	Images string
	UserID int
}

type PostIt struct {
	gorm.Model
	Body   string
	Images string
	Name   string
}

type CommentIt struct {
	gorm.Model
	Body   string
	Name   string
	PostID int
}

func FromDomain(do domain.Core) Post {
	return Post{
		Model:  gorm.Model{ID: do.ID},
		Body:   do.Body,
		Images: do.Images,
		UserID: do.UserID,
	}
}

func ToDomain(p PostIt) domain.Cores {
	return domain.Cores{
		Model:  gorm.Model{ID: p.ID, CreatedAt: p.CreatedAt},
		Body:   p.Body,
		Images: p.Images,
		Name:   p.Name,
	}
}

func ToDomainArray(ap []Post) []domain.Core {
	var res []domain.Core
	for _, val := range ap {
		res = append(res, domain.Core{Model: gorm.Model{ID: val.ID}, Body: val.Body, Images: val.Images, UserID: val.UserID})
	}

	return res
}

func ToDomainArrayIt(ai []PostIt) []domain.Cores {
	var res []domain.Cores
	for _, val := range ai {
		res = append(res, domain.Cores{Model: gorm.Model{ID: val.ID, CreatedAt: val.CreatedAt}, Body: val.Body, Images: val.Images, Name: val.Name})
	}

	return res
}

func ToDomainCommentIt(ci []CommentIt) []domain.Comes {
	var res []domain.Comes
	for _, val := range ci {
		res = append(res, domain.Comes{Model: gorm.Model{CreatedAt: val.CreatedAt}, Body: val.Body, Name: val.Name, PostID: val.PostID})
	}

	return res
}
