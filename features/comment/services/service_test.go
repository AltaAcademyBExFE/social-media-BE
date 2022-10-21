package services

import (
	"errors"
	dom "sosmed/features/comment/domain"
	rep "sosmed/features/comment/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var queryRepo = &rep.RepoQueryMock{Mock: mock.Mock{}}
var serviceRepo = CommentService{qry: queryRepo}

func TestCreate(t *testing.T) {
	repo := new(rep.RepoQueryMock)
	insertData := dom.Core{Body: "Testing Post.", PostID: 1, UserID: 1}
	returnData := dom.Cores{Body: "Testing Post."}

	t.Run("Success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)

		res, err := srv.Create(insertData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.Body, res.Body)
		//assert.Equal(t, returnData.Name, res.Name)
		//assert.Equal(t, returnData.CreatedAt, res.CreatedAt)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(dom.Core{}, errors.New("some problem on database")).Once()
		srv := New(repo)

		res, err := srv.Create(dom.Core{})
		assert.Equal(t, "", res.Body)
		assert.Equal(t, "", res.Name)
		assert.Error(t, err)
		assert.EqualError(t, err, "some problem on database")
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(rep.RepoQueryMock)
	//returnData := dom.Core{ID: uint(1), Email: "geger@gmail.com", Name: "Gerry", Phone: "1234", Address: "Malang"}

	t.Run("Success", func(t *testing.T) {
		repo.On("Del", 1).Return(nil).Once()

		srv := New(repo)

		err := srv.Delete(1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Out of Index", func(t *testing.T) {
		repo.On("Del", 2).Return(errors.New("error out of index")).Once()

		srv := New(repo)

		err := srv.Delete(2)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}
