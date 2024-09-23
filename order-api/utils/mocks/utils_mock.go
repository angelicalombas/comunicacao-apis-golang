package mocks

import (
	"github.com/stretchr/testify/mock"
)

type UtilsMock struct {
	mock.Mock
}

func (m *UtilsMock) CheckUserExists(userID int) (bool, error) {
	args := m.Called(userID)
	return args.Bool(0), args.Error(1)
}
