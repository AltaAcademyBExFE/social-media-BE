package repository

import (
	"sosmed/features/comment/domain"
	"time"

	"gorm.io/gorm"
)

type RepoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &RepoQuery{
		db: dbConn,
	}
}

func (rq *RepoQuery) Insert(newComment domain.Core) (domain.Cores, error) {
	var resQry CommentIt
	if err := rq.db.Exec("INSERT INTO comments (id, created_at, updated_at, deleted_at, body, post_id, user_id) values (?,?,?,?,?,?,?)",
		nil, time.Now(), nil, nil, newComment.Body, newComment.PostID, newComment.UserID).Error; err != nil {
		return domain.Cores{}, err
	}
	if er := rq.db.Table("comments").Select("comments.created_at", "comments.body", "users.name").Joins("join users on users.id=comments.user_id").Where("comments.body = ? AND comments.user_id = ?", newComment.Body, newComment.UserID).Model(&CommentIt{}).Find(&resQry).Error; er != nil {
		return domain.Cores{}, er
	}
	res := ToDomain(resQry)
	return res, nil
}

func (rq *RepoQuery) Del(ID int) error {
	var resQry Comment
	if err := rq.db.Where("id = ?", ID).Delete(&resQry).Error; err != nil {
		return err
	}
	return nil
}
