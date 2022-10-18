package repository

import (
	"sosmed/features/user/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) GetMyUser() (domain.UserCore, error) {
	var resQuery User
	if err := rq.db.First(&resQuery, "id = ?", resQuery.ID).Error; err != nil {
		log.Error("error on get my user", err.Error())
		return domain.UserCore{}, err
	}
	res := ToDomain(resQuery)
	return res, nil
}

func (rq *repoQuery) Update(updatedUser domain.UserCore) (domain.UserCore, error) {
	var cnv User = FromDomain(updatedUser)
	if err := rq.db.Where("id = ?", updatedUser.ID).Save(&cnv).Error; err != nil {
		log.Error("error on update user", err.Error())
		return domain.UserCore{}, err
	}
	if rq.db.RowsAffected == 0 {
		log.Info("content not updated")
		return domain.UserCore{}, nil
	}
	updatedUser = ToDomain(cnv)
	return updatedUser, nil
}

func (rq *repoQuery) Delete(deletedUser domain.UserCore) (domain.UserCore, error) {
	var cnv User = FromDomain(deletedUser)
	if err := rq.db.Where("id = ?", deletedUser.ID).Delete(&cnv).Error; err != nil {
		log.Error("error on deleting user", err.Error())
		return domain.UserCore{}, err
	}
	res := ToDomain(cnv)
	return res, nil
}

func (rq *repoQuery) GetByEmail(email string) (domain.UserCore, error) {
	var resQuery User
	if err := rq.db.First(&resQuery, "email = ?", email).Error; err != nil {
		log.Error("error on get by email", err.Error())
		return domain.UserCore{}, err
	}
	res := ToDomain(resQuery)
	return res, nil
}

func (rq *repoQuery) AddUser(newUser domain.UserCore) (domain.UserCore, error) {
	var cnv User = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on adding user", err.Error())
		return domain.UserCore{}, err
	}
	newUser = ToDomain(cnv)
	return newUser, nil
}

func (rq *repoQuery) GetUser(existUser domain.UserCore) (domain.UserCore, error) {
	var resQuery User
	if err := rq.db.First(&resQuery, "email = ?", existUser.Email).Error; err != nil {
		log.Error("error on get user login", err.Error())
		return domain.UserCore{}, nil
	}
	res := ToDomain(resQuery)
	return res, nil
}
