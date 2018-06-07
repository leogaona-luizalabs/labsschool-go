package mocks

import (
	"github.com/luizalabs/go-training/services/routes"
	"github.com/stretchr/testify/mock"
)

// RoutesMock ...
type RoutesMock struct {
	mock.Mock
}

// List ...
func (m *RoutesMock) List(limit int, offset int) (routes.Routes, error) {
	args := m.Called(limit, offset)

	if err := args.Get(1); err != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(routes.Routes), nil
}
