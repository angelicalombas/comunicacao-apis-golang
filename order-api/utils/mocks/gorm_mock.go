package mocks

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type GormDBMock struct {
	mock.Mock
}

func (m *GormDBMock) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(append([]interface{}{dest}, conds...)...)
	return args.Get(0).(*gorm.DB)
}

func (m *GormDBMock) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(append([]interface{}{dest}, conds...)...)
	return args.Get(0).(*gorm.DB)
}

func (m *GormDBMock) Where(query interface{}, args ...interface{}) *gorm.DB {
	calledArgs := m.Called(append([]interface{}{query}, args...)...)
	return calledArgs.Get(0).(*gorm.DB)
}

func (m *GormDBMock) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *GormDBMock) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *GormDBMock) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(append([]interface{}{value}, conds...)...)
	return args.Get(0).(*gorm.DB)
}
