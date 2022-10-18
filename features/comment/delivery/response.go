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
	Body      string    `json:"body"`
	Name      string    `json:"user"`
	CreatedAt time.Time `json:"create"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "comment":
		cnv := core.(domain.Cores)
		res = Responses{Body: cnv.Body, Name: cnv.Name, CreatedAt: cnv.CreatedAt}
	}
	return res
}
