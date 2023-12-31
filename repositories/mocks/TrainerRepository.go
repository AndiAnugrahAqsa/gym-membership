// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "ourgym/models"

	mock "github.com/stretchr/testify/mock"
)

// TrainerRepository is an autogenerated mock type for the TrainerRepository type
type TrainerRepository struct {
	mock.Mock
}

// CountTrainer provides a mock function with given fields:
func (_m *TrainerRepository) CountTrainer() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetAll provides a mock function with given fields: name
func (_m *TrainerRepository) GetAll(name string) []models.Trainer {
	ret := _m.Called(name)

	var r0 []models.Trainer
	if rf, ok := ret.Get(0).(func(string) []models.Trainer); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Trainer)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *TrainerRepository) GetByID(id string) models.Trainer {
	ret := _m.Called(id)

	var r0 models.Trainer
	if rf, ok := ret.Get(0).(func(string) models.Trainer); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Trainer)
	}

	return r0
}

type mockConstructorTestingTNewTrainerRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTrainerRepository creates a new instance of TrainerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTrainerRepository(t mockConstructorTestingTNewTrainerRepository) *TrainerRepository {
	mock := &TrainerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
