package delivery

import "sosmed/features/comment/domain"

type CommentFormat struct {
	Body   string `json:"body" form:"body"`
	PostID int    `json:"post" form:"post"`
	UserID int    `json:"user" form:"user"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case CommentFormat:
		cnv := i.(CommentFormat)
		return domain.Core{Body: cnv.Body, PostID: cnv.PostID, UserID: cnv.UserID}
	}
	return domain.Core{}
}
