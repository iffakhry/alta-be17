package service

import (
	"be17/cleanarch/features/user"
	"be17/cleanarch/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	userDataLayer := new(mocks.UserData)
	returnData := []user.Core{
		{Id: 1, Name: "Andi", Phone: "0812345", Email: "andi@mail.com", Password: "abcd"},
	}
	t.Run("test case success get all data", func(t *testing.T) {
		userDataLayer.On("SelectAll").Return(returnData, nil).Once()
		srv := New(userDataLayer)
		response, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		assert.Equal(t, returnData[0].Email, response[0].Email)
		userDataLayer.AssertExpectations(t)
	})

	t.Run("test case failed get all data", func(t *testing.T) {
		userDataLayer.On("SelectAll").Return(nil, errors.New("error read data from db")).Once()
		srv := New(userDataLayer)
		response, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, response)
		userDataLayer.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	userDataLayer := new(mocks.UserData)

	t.Run("test case success insert data", func(t *testing.T) {
		insertData := user.Core{Name: "Andi", Phone: "0812345", Email: "andi@mail.com", Password: "abcd"}
		userDataLayer.On("Insert", insertData).Return(nil).Once()
		srv := New(userDataLayer)
		err := srv.Create(insertData)
		assert.Nil(t, err)
		userDataLayer.AssertExpectations(t)
	})

	t.Run("test case failed insert data to db", func(t *testing.T) {
		insertData := user.Core{Name: "Andi", Phone: "0812345", Email: "andi@mail.com", Password: "abcd"}
		userDataLayer.On("Insert", insertData).Return(errors.New("insert failed")).Once()
		srv := New(userDataLayer)
		err := srv.Create(insertData)
		assert.NotNil(t, err)
		userDataLayer.AssertExpectations(t)
	})

	t.Run("test case failed insert data, validation error", func(t *testing.T) {
		insertData := user.Core{Name: "Andi", Phone: "0812345"}

		srv := New(userDataLayer)
		err := srv.Create(insertData)
		assert.NotNil(t, err)
		userDataLayer.AssertExpectations(t)
	})
}
