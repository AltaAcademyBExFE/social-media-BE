package services

import (
	"errors"
	"sosmed/features/post/domain"
	"strings"

	"github.com/labstack/gommon/log"
)

type PostService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &PostService{
		qry: repo,
	}
}

func (ps *PostService) ShowAll() ([]domain.Cores, []domain.Comes, error) {
	res, rel, err := ps.qry.Show()
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, nil, errors.New("Database Error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, nil, errors.New("No Data")
		}
	}

	if len(res) == 0 {
		log.Info("No Data")
		return nil, nil, errors.New("No Data")
	}
	return res, rel, nil
}

func (ps *PostService) ShowMy(ID int) ([]domain.Cores, []domain.Comes, error) {
	res, rel, err := ps.qry.My(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, nil, errors.New("Database Error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, nil, errors.New("No Data")
		}
	}

	return res, rel, nil
}

func (ps *PostService) ShowSpesific(ID int) ([]domain.Cores, []domain.Comes, error) {
	res, rel, err := ps.qry.Spesific(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, nil, errors.New("Database Error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, nil, errors.New("No Data")
		}
	}

	return res, rel, nil
}

func (ps *PostService) Create(newPost domain.Core) (domain.Cores, error) {
	res, err := ps.qry.Insert(newPost)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Cores{}, errors.New("Rejected from Database")
		}

		return domain.Cores{}, errors.New("Some Problem on Database")
	}

	return res, nil
}

func (ps *PostService) Edit(ID int, updatePost domain.Core) (domain.Cores, error) {
	res, err := ps.qry.Update(ID, updatePost)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Cores{}, errors.New("Rejected from Database")
		}

		return domain.Cores{}, errors.New("Some Problem on Database")
	}

	return res, nil
}

func (ps *PostService) Delete(ID int) error {
	err := ps.qry.Del(ID)
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
