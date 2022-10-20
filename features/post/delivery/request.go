package delivery

import (
	"sosmed/features/post/domain"
)

type PostingFormat struct {
	Body   string `json:"body" form:"body"`
	Images string `json:"img" form:"img"`
	UserID int    `json:"user" form:"user"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case PostingFormat:
		cnv := i.(PostingFormat)
		return domain.Core{Body: cnv.Body, Images: cnv.Images, UserID: cnv.UserID}
	}
	return domain.Core{}
}
