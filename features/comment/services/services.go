package services

import (
	"errors"
	"sosmed/features/comment/domain"
	"strings"

	"github.com/labstack/gommon/log"
)

type CommentService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &CommentService{
		qry: repo,
	}
}

func (cs *CommentService) Create(newComment domain.Core) (domain.Core, error) {
	res, err := cs.qry.Insert(newComment)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("Rejected from Database")
		}

		return domain.Core{}, errors.New("Some Problem on Database")
	}

	return res, nil
}

func (cs *CommentService) Delete(ID int) error {
	err := cs.qry.Del(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return errors.New("Database Error")
		} else if strings.Contains(err.Error(), "found") {
			return errors.New("No Data")
		}
	}
	return nil
}
