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

func FromDomain(do domain.Core) Post {
	return Post{
		Model:  gorm.Model{ID: do.ID},
		Body:   do.Body,
		Images: do.Images,
		UserID: do.UserID,
	}
}

func ToDomain(p Post) domain.Core {
	return domain.Core{
		Model:  gorm.Model{ID: p.ID},
		Body:   p.Body,
		Images: p.Images,
		UserID: p.UserID,
	}
}

func ToDomainArray(ap []Post) []domain.Core {
	var res []domain.Core
	for _, val := range ap {
		res = append(res, domain.Core{Model: gorm.Model{ID: val.ID}, Body: val.Body, Images: val.Images, UserID: val.UserID})
	}

	return res
}
