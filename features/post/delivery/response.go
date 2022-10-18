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
	Name      string    `json:"user"`
	CreatedAt time.Time `json:"create"`
}

type Response struct {
	Body      string    `json:"body"`
	Images    string    `json:"img"`
	Name      string    `json:"user"`
	CreatedAt time.Time `json:"create"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "post":
		cnv := core.(domain.Cores)
		res = Response{Body: cnv.Body, Images: cnv.Images, Name: cnv.Name, CreatedAt: cnv.CreatedAt}
	case "all":
		var arr []Responses
		cnv := core.([]domain.Cores)
		for _, val := range cnv {
			arr = append(arr, Responses{ID: val.ID, Body: val.Body, Images: val.Images, Name: val.Name, CreatedAt: val.CreatedAt})
		}
		res = arr
	}

	return res
}
