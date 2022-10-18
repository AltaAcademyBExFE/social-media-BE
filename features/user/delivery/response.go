package delivery

import "sosmed/features/user/domain"

type RegisterResponse struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "user":
		cnv := core.(domain.UserCore)
		res = RegisterResponse{Name: cnv.Name, Email: cnv.Email, Phone: cnv.Phone, Address: cnv.Address}
	case "login":
		cnv := core.(domain.UserCore)
		res = LoginResponse{Email: cnv.Email, Token: cnv.Token}
	}
	return res
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Message": msg,
		"Data":    data,
	}
}

func FailResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"Message": msg,
	}
}
