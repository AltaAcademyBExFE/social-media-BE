package services

import (
	"errors"
	"sosmed/features/user/domain"
	"sosmed/features/user/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMyProfile(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("success", func(t *testing.T) {
		repo.On("GetMyUser", mock.Anything).Return(domain.UserCore{ID: uint(1), Name: "Dany", Email: "dany@gmail.com", Phone: "08123", Password: "dany123"}, nil).Once()
		srv := New(repo)
		res, err := srv.MyProfile(uint(1))
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get User", func(t *testing.T) {
		repo.On("GetMyUser", mock.Anything).Return(domain.UserCore{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.MyProfile(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateProfile(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Update User", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(domain.UserCore{ID: uint(1), Name: "Dany", Email: "dany@gmail.com", Phone: "08123", Address: "Malang", Password: "dany123"}, nil).Once()
		srv := New(repo)
		input := domain.UserCore{ID: uint(1), Name: "Dany", Email: "dany@gmail.com", Phone: "08123", Address: "Malang", Password: "dany123"}
		res, err := srv.UpdateProfile(input, 1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Update User", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(domain.UserCore{}, errors.New("some problem on database")).Once()
		srv := New(repo)
		var input domain.UserCore
		res, err := srv.UpdateProfile(input, 1)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDeactivate(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Delete User", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.UserCore{ID: uint(1), Name: "Dany", Email: "dany@gmail.com", Phone: "08123", Address: "Malang", Password: "dany123"}, nil).Once()
		srv := New(repo)
		res, err := srv.Deactivate(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete User", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.UserCore{}, errors.New("database error")).Once()
		srv := New(repo)
		res, err := srv.Deactivate(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Add User", func(t *testing.T) {
		// repo.On("GetByUsername", mock.Anything).Return(domain.Core{}, 0)
		repo.On("AddUser", mock.Anything).Return(domain.UserCore{ID: uint(1), Name: "Dany", Email: "dany@gmail.com", Phone: "08123", Address: "Malang", Password: "dany123"}, nil).Once()
		srv := New(repo)
		input := domain.UserCore{ID: uint(1), Name: "Dany", Email: "dany@gmail.com", Phone: "08123", Address: "Malang", Password: "dany123"}
		res, err := srv.Register(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Add User", func(t *testing.T) {
		repo.On("AddUser", mock.Anything).Return(domain.UserCore{}, errors.New("some problem on database")).Once()
		srv := New(repo)
		res, err := srv.Register(domain.UserCore{})
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(domain.UserCore{Password: "$2a$10$szpOHiZl0Uvv.Wr1hTAwKeSbTb2E2igPNzPHqW.C0u5xMmLRomaYS "}, nil).Once()
		srv := New(repo)
		input := domain.UserCore{Email: "dany@gmail.com", Password: "dany123"}
		res, err := srv.Login(input)
		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Wrong username Login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(domain.UserCore{Password: "asgfasg"}, errors.New("no data")).Once()
		srv := New(repo)
		input := domain.UserCore{Email: "dany@gmail.com", Password: "dany123"}
		res, err := srv.Login(input)
		assert.Empty(t, res)
		assert.EqualError(t, err, "no data")
		repo.AssertExpectations(t)
	})
}
