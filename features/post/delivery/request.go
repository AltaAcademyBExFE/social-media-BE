package delivery

import (
	"sosmed/features/post/domain"

	"gorm.io/gorm"
)

type PostingFormat struct {
	Body   string `json:"body" form:"body"`
	Images string `json:"img" form:"img"`
	UserID int    `json:"user" form:"user"`
}

type EditFormat struct {
	ID     uint   `json:"id" form:"id"`
	Body   string `json:"body" form:"body"`
	Images string `json:"img" form:"img"`
	UserID int    `json:"user" form:"user"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case PostingFormat:
		cnv := i.(PostingFormat)
		return domain.Core{Body: cnv.Body, Images: cnv.Images, UserID: cnv.UserID}
	case EditFormat:
		cnv := i.(EditFormat)
		return domain.Core{Model: gorm.Model{ID: cnv.ID}, Body: cnv.Body, Images: cnv.Images, UserID: cnv.UserID}
	}
	return domain.Core{}
}
