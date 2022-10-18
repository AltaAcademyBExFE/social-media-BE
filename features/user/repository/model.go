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
}

func FromDomain(du domain.UserCore) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Name:     du.Name,
		Email:    du.Email,
		Phone:    du.Phone,
		Address:  du.Address,
		Password: du.Password,
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
	}
}

func ToDomainArray(ua []User) []domain.UserCore {
	var res []domain.UserCore
	for _, val := range ua {
		res = append(res, domain.UserCore{
			ID:       val.ID,
			Name:     val.Name,
			Email:    val.Email,
			Phone:    val.Phone,
			Address:  val.Address,
			Password: val.Password,
		})
	}
	return res
}
