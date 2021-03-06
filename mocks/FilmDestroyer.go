// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	context "context"

	film "github.com/migalpha/kentech-films"
	mock "github.com/stretchr/testify/mock"
)

// FilmDestroyer is an autogenerated mock type for the FilmDestroyer type
type FilmDestroyer struct {
	mock.Mock
}

// Destroy provides a mock function with given fields: _a0, _a1
func (_m *FilmDestroyer) Destroy(_a0 context.Context, _a1 film.FilmID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, film.FilmID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
