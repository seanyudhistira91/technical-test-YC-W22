package mock_repositories

import (
	"context"

	"github.com/seanyudhistira91/technical-test-YC-W22/entity"
	"github.com/seanyudhistira91/technical-test-YC-W22/repository"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

// func NewMockRepository() *MockRepository {
// 	return &MockRepository{}
// }

// implement interface
var _ repository.Repository = &MockRepository{}

func (m *MockRepository) MockCreate(err error) {
	m.Mock.On("Create", mock.Anything, mock.Anything).
		Return(err)
}

func (m *MockRepository) Create(
	ctx context.Context,
	e entity.Users,
) error {
	args := m.Called(ctx, e)
	var err error

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}

func (m *MockRepository) MockGetByPhoneNumber(e *entity.Users, err error) {
	m.Mock.On("GetByPhoneNumber", mock.Anything, mock.Anything).
		Return(e, err)
}

func (m *MockRepository) GetByPhoneNumber(
	ctx context.Context,
	phoneNumber string,
) (*entity.Users, error) {
	args := m.Called(ctx, phoneNumber)
	var err error
	var e *entity.Users

	if n, ok := args.Get(0).(*entity.Users); ok {
		e = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return e, err
}
