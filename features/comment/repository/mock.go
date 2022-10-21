package repository

import (
	"sosmed/features/comment/domain"

	"github.com/stretchr/testify/mock"
)

type RepoQueryMock struct {
	mock.Mock
}

func (rq *RepoQueryMock) Insert(newUser domain.Core) (domain.Cores, error) {
	ret := rq.Called(newUser)

	var r0 domain.Cores
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Cores); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(domain.Cores)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (rq *RepoQueryMock) Del(ID int) error {
	ret := rq.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(int(ID))
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
