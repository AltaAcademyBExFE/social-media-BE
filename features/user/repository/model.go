package repository

import (
	"sosmed/features/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Phone    string
	Address  string
	Password string
	Token    string
}

func FromDomain(du domain.UserCore) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Name:     du.Name,
		Email:    du.Email,
		Phone:    du.Phone,
		Address:  du.Address,
		Password: du.Password,
		Token:    du.Token,
	}
}

func ToDomain(u User) domain.UserCore {
	return domain.UserCore{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Phone:    u.Phone,
		Address:  u.Address,
		Password: u.Password,
		Token:    u.Token,
	}
}
