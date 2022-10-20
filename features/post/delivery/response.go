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
	ID        uint        `json:"id"`
	Body      string      `json:"body"`
	Images    string      `json:"img"`
	Name      string      `json:"user"`
	CreatedAt time.Time   `json:"create"`
	Comments  []Responser `json:"comments"`
}

type Response struct {
	Body      string    `json:"body"`
	Images    string    `json:"img"`
	Name      string    `json:"user"`
	CreatedAt time.Time `json:"create"`
}

type Responser struct {
	Body      string    `json:"body"`
	Name      string    `json:"user"`
	CreatedAt time.Time `json:"create"`
	PostID    int       `json:"id_post"`
}

func ToResponse(core interface{}, come interface{}, code string) interface{} {
	var res interface{}
	//var rel interface{}
	switch code {
	case "post":
		cnv := core.(domain.Cores)
		res = Response{Body: cnv.Body, Images: cnv.Images, Name: cnv.Name, CreatedAt: cnv.CreatedAt}
	case "all":
		var arr []Responses
		cnv := core.([]domain.Cores)
		for _, val := range cnv {

			var ar []Responser
			cnr := come.([]domain.Comes)
			for _, vol := range cnr {
				if vol.PostID == int(val.ID) && len(ar) < 3 {
					ar = append(ar, Responser{Body: vol.Body, Name: vol.Name, CreatedAt: vol.CreatedAt, PostID: vol.PostID})
				}
			}

			arr = append(arr, Responses{ID: val.ID, Body: val.Body, Images: val.Images, Name: val.Name, CreatedAt: val.CreatedAt, Comments: ar})

		}
		res = arr
	case "ally":
		var arr []Responses
		cnv := core.([]domain.Cores)
		for _, val := range cnv {

			var ar []Responser
			cnr := come.([]domain.Comes)
			for _, vol := range cnr {
				if vol.PostID == int(val.ID) {
					ar = append(ar, Responser{Body: vol.Body, Name: vol.Name, CreatedAt: vol.CreatedAt, PostID: vol.PostID})
				}
			}

			arr = append(arr, Responses{ID: val.ID, Body: val.Body, Images: val.Images, Name: val.Name, CreatedAt: val.CreatedAt, Comments: ar})

		}
		res = arr
	}

	return res
}

func ToResponser(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "comments":
		var arr []Responser
		cnv := core.([]domain.Comes)
		for _, val := range cnv {
			arr = append(arr, Responser{Body: val.Body, Name: val.Name, CreatedAt: val.CreatedAt})
		}
		res = arr
	}
	return res
}
