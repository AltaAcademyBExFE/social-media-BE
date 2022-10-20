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

func (rq *repoQuery) Show() ([]domain.Cores, []domain.Comes, error) {
	var resQry []PostIt
	var comQry []CommentIt
	if err := rq.db.Table("posts").Select("posts.id", "posts.created_at", "posts.body", "posts.images", "users.name").Joins("join users on users.id=posts.user_id").Model(&PostIt{}).Find(&resQry).Error; err != nil {
		return nil, nil, err
	}
	if er := rq.db.Table("comments").Select("comments.created_at", "comments.body", "users.name", "comments.post_id").Joins("join users on users.id=comments.user_id").Model(&CommentIt{}).Find(&comQry).Error; er != nil {
		return nil, nil, er
	}
	rel := ToDomainCommentIt(comQry)
	res := ToDomainArrayIt(resQry)
	return res, rel, nil
}

func (rq *repoQuery) My(ID int) ([]domain.Cores, []domain.Comes, error) {
	var resQry []PostIt
	var comQry []CommentIt
	if err := rq.db.Table("posts").Select("posts.id", "posts.created_at", "posts.body", "posts.images", "users.name").Joins("join users on users.id=posts.user_id").Where("users.id = ?", ID).Model(&PostIt{}).Find(&resQry).Error; err != nil {
		return nil, nil, err
	}
	if er := rq.db.Table("comments").Select("comments.created_at", "comments.body", "users.name", "comments.post_id").Joins("join users on users.id=comments.user_id").Joins("join posts on posts.id=comments.post_id").Where("posts.user_id = ?", ID).Model(&CommentIt{}).Find(&comQry).Error; er != nil {
		return nil, nil, er
	}
	rel := ToDomainCommentIt(comQry)
	res := ToDomainArrayIt(resQry)
	return res, rel, nil
}

func (rq *repoQuery) Spesific(ID int) ([]domain.Cores, []domain.Comes, error) {
	var resQry []PostIt
	var comQry []CommentIt
	if err := rq.db.Table("posts").Select("posts.id", "posts.created_at", "posts.body", "posts.images", "users.name").Joins("join users on users.id=posts.user_id").Where("posts.id = ?", ID).Model(&PostIt{}).Find(&resQry).Error; err != nil {
		return nil, nil, err
	}
	if er := rq.db.Table("comments").Select("comments.created_at", "comments.body", "users.name", "comments.post_id").Joins("join users on users.id=comments.user_id").Joins("join posts on posts.id=comments.post_id").Where("posts.id = ?", ID).Model(&CommentIt{}).Find(&comQry).Error; er != nil {
		return nil, nil, er
	}
	rel := ToDomainCommentIt(comQry)
	res := ToDomainArrayIt(resQry)
	return res, rel, nil
}

func (rq *repoQuery) Insert(newPost domain.Core) (domain.Cores, error) {
	var resQry PostIt
	if err := rq.db.Exec("INSERT INTO posts (id, created_at, updated_at, deleted_at, body, images, user_id) values (?,?,?,?,?,?,?)",
		nil, time.Now(), nil, nil, newPost.Body, newPost.Images, newPost.UserID).Error; err != nil {
		return domain.Cores{}, err
	}
	if er := rq.db.Table("posts").Select("posts.created_at", "posts.body", "posts.images", "users.name").Joins("join users on users.id=posts.user_id").Where("posts.body = ? AND posts.images = ? AND posts.user_id = ?", newPost.Body, newPost.Images, newPost.UserID).Model(&PostIt{}).Find(&resQry).Error; er != nil {
		return domain.Cores{}, er
	}
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) Update(ID int, updatePost domain.Core) (domain.Cores, error) {
	var resQry PostIt
	if err := rq.db.Exec("UPDATE posts SET updated_at = ?, body = ?, images = ? WHERE id = ?",
		time.Now(), updatePost.Body, updatePost.Images, ID).Error; err != nil {
		return domain.Cores{}, err
	}
	if er := rq.db.Table("posts").Select("posts.created_at", "posts.body", "posts.images", "users.name").Joins("join users on users.id=posts.user_id").Where("posts.id = ?", ID).Model(&PostIt{}).Find(&resQry).Error; er != nil {
		return domain.Cores{}, er
	}
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) Del(ID int) error {
	var resQry Post
	if err := rq.db.Where("id = ?", ID).Delete(&resQry).Error; err != nil {
		return err
	}
	return nil
}
