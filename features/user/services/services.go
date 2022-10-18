package services

import (
	"errors"
	"sosmed/features/user/domain"
	"strings"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{qry: repo}
}

func (us *userService) MyProfile() (domain.UserCore, error) {
	res, err := us.qry.GetMyUser()
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.UserCore{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.UserCore{}, errors.New("no data")
		}
	}
	return res, nil
}

func (us *userService) UpdateProfile(updatedUser domain.UserCore) (domain.UserCore, error) {
	res, err := us.qry.Update(updatedUser)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.UserCore{}, errors.New("rejected from database")
		}
		return domain.UserCore{}, errors.New("some problem on database")
	}
	return res, nil
}

func (us *userService) Deactivate(deletedUser domain.UserCore) (domain.UserCore, error) {
	res, err := us.qry.Delete(deletedUser)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.UserCore{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.UserCore{}, errors.New("no data")
		}
	}
	return res, nil
}

func (us *userService) ShowByEmail(email string) (domain.UserCore, error) {
	res, err := us.qry.GetByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.UserCore{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.UserCore{}, errors.New("no data")
		}
	}
	return res, nil
}

func (us *userService) Register(newUser domain.UserCore) (domain.UserCore, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error on bcrypt", err.Error())
		return domain.UserCore{}, errors.New("cannot encrypt password")
	}
	newUser.Password = string(generate)

	res, err := us.qry.AddUser(newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.UserCore{}, errors.New("rejected from database")
		}
		return domain.UserCore{}, errors.New("some problem on database")
	}
	return res, nil
}

func (us *userService) Login(existUser domain.UserCore) (domain.UserCore, error) {
	res, err := us.qry.GetUser(existUser)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.UserCore{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.UserCore{}, errors.New("no data")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(existUser.Password))
	if err != nil {
		return domain.UserCore{}, errors.New("password not match")
	}

	return res, nil
}
