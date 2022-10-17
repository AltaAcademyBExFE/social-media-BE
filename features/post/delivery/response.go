package delivery

import (
	"sosmed/features/post/domain"
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
	Images    string    `json:"img"`
	UserID    int       `json:"user"`
	CreatedAt time.Time `json:"create"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "post":
		cnv := core.(domain.Core)
		res = Responses{ID: cnv.ID, Body: cnv.Body, Images: cnv.Images, UserID: cnv.UserID, CreatedAt: cnv.CreatedAt}
	case "all":
		var arr []Responses
		cnv := core.([]domain.Core)
		for _, val := range cnv {
			arr = append(arr, Responses{ID: val.ID, Body: val.Body, Images: val.Images, UserID: val.UserID, CreatedAt: val.CreatedAt})
		}
		res = arr
	}

	return res
}
