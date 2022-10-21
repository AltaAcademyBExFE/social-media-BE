package services

import (
	"errors"
	dom "sosmed/features/comment/domain"
	mockery "sosmed/features/comment/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var queryRepo = &mockery.Repository{Mock: mock.Mock{}}
var serviceRepo = CommentService{qry: queryRepo}

func TestCreate(t *testing.T) {
	repo := new(mockery.Repository)
	insertData := dom.Core{Body: "Testing Post.", PostID: 1, UserID: 1}
	returnData := dom.Cores{Body: "Testing Post.", Name: "Gerry"}

	t.Run("Success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)

		res, err := srv.Create(insertData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.Body, res.Body)
		assert.Equal(t, returnData.Name, res.Name)
	})

	t.Run("Error insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(dom.Cores{}, errors.New("some problem on database")).Once()
		srv := New(repo)

		res, err := srv.Create(dom.Core{})
		assert.Equal(t, "", res.Body)
		assert.Equal(t, "", res.Name)
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mockery.Repository)

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
