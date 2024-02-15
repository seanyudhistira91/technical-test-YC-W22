package service

import (
	"context"
	"testing"

	"errors"

	"github.com/seanyudhistira91/technical-test-YC-W22/entity"
	mock_entities "github.com/seanyudhistira91/technical-test-YC-W22/mocks/entities"
	mock_repositories "github.com/seanyudhistira91/technical-test-YC-W22/mocks/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockService struct {
	mock.Mock
}

type ServiceSuite struct {
	suite.Suite
	service  Service
	mockRepo mock_repositories.MockRepository

	userEntity entity.Users
}

func (s *ServiceSuite) SetupTest() {
	s.mockRepo = mock_repositories.MockRepository{}
	s.service = NewService(
		&s.mockRepo,
	)
	s.userEntity = mock_entities.MockUsers()

}

func (s *ServiceSuite) Test_CreateUser_Success() {
	ctx := context.Background()

	s.mockRepo.MockCreate(nil)
	err := s.service.CreateUser(ctx, s.userEntity)

	s.Nil(err)
}

func (s *ServiceSuite) Test_CreateUser_ServiceErr() {
	ctx := context.Background()

	s.mockRepo.MockCreate(errors.New("err"))
	err := s.service.CreateUser(ctx, s.userEntity)

	s.NotNil(err)
}

func (s *ServiceSuite) Test_Login_Success() {
	ctx := context.Background()

	s.userEntity = mock_entities.MockUsersWithHashPassword()
	s.mockRepo.MockGetByPhoneNumber(&s.userEntity, nil)
	user := entity.Users{
		PhoneNumber:  "0822222222",
		HashPassword: "testPassword",
	}
	err := s.service.Login(ctx, user)

	s.Nil(err)
}

func (s *ServiceSuite) Test_Login_ErrRepo() {
	ctx := context.Background()

	// s.userEntity = mock_entities.MockUsersWithHashPassword()
	s.mockRepo.MockGetByPhoneNumber(nil, errors.New("err"))
	user := entity.Users{
		PhoneNumber:  "0822222222",
		HashPassword: "testPassword",
	}
	err := s.service.Login(ctx, user)

	s.NotNil(err)
}

func (s *ServiceSuite) Test_Login_ErrMatchPassword() {
	ctx := context.Background()

	s.userEntity = mock_entities.MockUsersWithHashPassword()
	s.mockRepo.MockGetByPhoneNumber(&s.userEntity, nil)
	user := entity.Users{
		PhoneNumber:  "0822222222",
		HashPassword: "testPassword123",
	}
	err := s.service.Login(ctx, user)

	s.NotNil(err)
}
func Test_UserService(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
