package repository

import (
	"sosmed/features/comment/domain"
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

func (rq *repoQuery) Insert(newComment domain.Core) (domain.Core, error) {
	var cnv Comment
	cnv = FromDomain(newComment)
	if err := rq.db.Exec("INSERT INTO comments (id, created_at, updated_at, deleted_at, body, post_id, user_id) values (?,?,?,?,?,?,?)",
		nil, time.Now(), nil, nil, newComment.Body, newComment.PostID, newComment.UserID).Error; err != nil {
		return domain.Core{}, err
	}
	newComment = ToDomain(cnv)
	return newComment, nil
}

func (rq *repoQuery) Del(ID int) error {
	var resQry Comment
	if err := rq.db.Where("id = ?", ID).Delete(&resQry).Error; err != nil {
		return err
	}
	return nil
}
