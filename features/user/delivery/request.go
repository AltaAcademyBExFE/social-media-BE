package delivery

import "sosmed/features/user/domain"

type UserFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Password string `json:"password" form:"password"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type GetUserFormat struct {
	Email string `json:"email" form:"email"`
}
type DeleteFormat struct {
	Name string `json:"name" form:"name"`
}

func ToDomain(i interface{}) domain.UserCore {
	switch i.(type) {
	case UserFormat:
		cnv := i.(UserFormat)
		return domain.UserCore{Name: cnv.Name, Email: cnv.Email, Phone: cnv.Phone, Address: cnv.Address, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.UserCore{Email: cnv.Email, Password: cnv.Password}
	case GetUserFormat:
		cnv := i.(GetUserFormat)
		return domain.UserCore{Email: cnv.Email}
	case DeleteFormat:
		cnv := i.(DeleteFormat)
		return domain.UserCore{Name: cnv.Name}
	}
	return domain.UserCore{}
}
