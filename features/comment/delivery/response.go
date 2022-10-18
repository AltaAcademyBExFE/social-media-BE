package delivery

import (
	"sosmed/features/comment/domain"
	"time"
)

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type Responses struct {
	ID        uint      `json:"id"`
	Body      string    `json:"body"`
	PostID    int       `json:"post"`
	UserID    int       `json:"user"`
	CreatedAt time.Time `json:"create"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "comment":
		cnv := core.(domain.Core)
		res = Responses{ID: cnv.ID, Body: cnv.Body, PostID: cnv.PostID, UserID: cnv.UserID, CreatedAt: cnv.CreatedAt}
	}

	return res
}
