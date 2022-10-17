package repository

import (
	"sosmed/features/post/domain"
	"time"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) Show() ([]domain.Core, error) {
	var resQry []Post
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) My(ID int) ([]domain.Core, error) {
	var resQry []Post
	if err := rq.db.Find(&resQry, "user_id = ?", ID).Error; err != nil {
		return nil, err
	}
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) Insert(newPost domain.Core) (domain.Core, error) {
	var cnv Post
	cnv = FromDomain(newPost)
	if err := rq.db.Exec("INSERT INTO posts (id, created_at, updated_at, deleted_at, body, images, user_id) values (?,?,?,?,?,?,?)",
		nil, time.Now(), nil, nil, newPost.Body, newPost.Images, newPost.UserID).Error; err != nil {
		return domain.Core{}, err
	}
	newPost = ToDomain(cnv)
	return newPost, nil
}

func (rq *repoQuery) Update(updatePost domain.Core) (domain.Core, error) {
	var cnv Post
	cnv = FromDomain(updatePost)
	if err := rq.db.Exec("UPDATE posts SET updated_at = ?, body = ?, images = ? WHERE id = ?",
		time.Now(), cnv.Body, cnv.Images, cnv.ID).Error; err != nil {
		return domain.Core{}, err
	}
	updatePost = ToDomain(cnv)
	return updatePost, nil
}

func (rq *repoQuery) Del(ID int) error {
	var resQry Post
	if err := rq.db.Where("id = ?", ID).Delete(&resQry).Error; err != nil {
		return err
	}
	return nil
}
