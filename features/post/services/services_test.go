package services

import (
	"errors"
	dom "sosmed/features/post/domain"
	mockery "sosmed/features/post/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var queryRepo = &mockery.Repository{Mock: mock.Mock{}}
var serviceRepo = PostService{qry: queryRepo}

func TestShowAll(t *testing.T) {
	repo := new(mockery.Repository)
	returnPost := []dom.Cores{{Body: "Testing Post.", Images: "Photos-1.jpg", Name: "Gerry"}}
	returnCom := []dom.Comes{{Body: "Testing Post.", Name: "Gerry", PostID: 1}}

	t.Run("Success", func(t *testing.T) {
		repo.On("Show").Return(returnPost, returnCom, nil).Once()

		srv := New(repo)

		res, rel, err := srv.ShowAll()
		assert.NoError(t, err)
		assert.Equal(t, returnPost[0].ID, res[0].ID)
		assert.Equal(t, returnCom[0].ID, rel[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		repo.On("Show").Return(nil, nil, errors.New("data not found")).Once()

		srv := New(repo)

		res, rel, err := srv.ShowAll()
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Nil(t, rel)
		repo.AssertExpectations(t)
	})
}

func TestShowMy(t *testing.T) {
	repo := new(mockery.Repository)
	returnPost := []dom.Cores{{Body: "Testing Post.", Images: "Photos-1.jpg", Name: "Gerry"}}
	returnCom := []dom.Comes{{Body: "Testing Post.", Name: "Gerry", PostID: 1}}

	t.Run("Success", func(t *testing.T) {
		repo.On("My", 1).Return(returnPost, returnCom, nil).Once()

		srv := New(repo)

		res, rel, err := srv.ShowMy(1)
		assert.NoError(t, err)
		assert.Equal(t, returnPost[0].ID, res[0].ID)
		assert.Equal(t, returnCom[0].ID, rel[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		repo.On("My", 1).Return(nil, nil, errors.New("data not found")).Once()

		srv := New(repo)

		res, rel, err := srv.ShowMy(1)
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Nil(t, rel)
		repo.AssertExpectations(t)
	})
}

func TestShowSpesific(t *testing.T) {
	repo := new(mockery.Repository)
	returnPost := []dom.Cores{{Body: "Testing Post.", Images: "Photos-1.jpg", Name: "Gerry"}}
	returnCom := []dom.Comes{{Body: "Testing Post.", Name: "Gerry", PostID: 1}}

	t.Run("Success", func(t *testing.T) {
		repo.On("Spesific", 1).Return(returnPost, returnCom, nil).Once()

		srv := New(repo)

		res, rel, err := srv.ShowSpesific(1)
		assert.NoError(t, err)
		assert.Equal(t, returnPost[0].ID, res[0].ID)
		assert.Equal(t, returnCom[0].ID, rel[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		repo.On("Spesific", 1).Return(nil, nil, errors.New("data not found")).Once()

		srv := New(repo)

		res, rel, err := srv.ShowSpesific(1)
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Nil(t, rel)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mockery.Repository)
	insertData := dom.Core{Body: "Testing Post.", Images: "Photos-1.jpg", UserID: 1}
	returnData := dom.Cores{Body: "Testing Post.", Images: "Photos-1.jpg", Name: "Gerry"}

	t.Run("Success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)

		res, err := srv.Create(insertData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.Body, res.Body)
		assert.Equal(t, returnData.Images, res.Images)
		assert.Equal(t, returnData.Name, res.Name)
	})

	t.Run("Error insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(dom.Cores{}, errors.New("Some Problem on Database")).Once()
		srv := New(repo)

		res, err := srv.Create(dom.Core{})
		assert.Equal(t, "", res.Body)
		assert.Equal(t, "", res.Images)
		assert.Equal(t, "", res.Name)
		assert.Error(t, err)
		assert.EqualError(t, err, "Some Problem on Database")
	})
}

func TestEdit(t *testing.T) {
	repo := new(mockery.Repository)
	updateData := dom.Core{Body: "Testing Update Post.", Images: "Photos-1.jpg", UserID: 1}
	returnData := dom.Cores{Body: "Testing Update Post.", Images: "Photos-1.jpg", Name: "Gerry"}

	t.Run("Success", func(t *testing.T) {
		repo.On("Update", 1, mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)

		res, err := srv.Edit(1, updateData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.Body, res.Body)
		assert.Equal(t, returnData.Images, res.Images)
		assert.Equal(t, returnData.Name, res.Name)
	})

	t.Run("Error insert", func(t *testing.T) {
		repo.On("Update", 1, mock.Anything).Return(dom.Cores{}, errors.New("Some Problem on Database")).Once()
		srv := New(repo)

		res, err := srv.Edit(1, dom.Core{})
		assert.Equal(t, "", res.Body)
		assert.Equal(t, "", res.Images)
		assert.Equal(t, "", res.Name)
		assert.Error(t, err)
		assert.EqualError(t, err, "Some Problem on Database")
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
