// service/service.go
package service

import (
	"context"

	"github.com/seanyudhistira91/technical-test-YC-W22/entity"
	"github.com/seanyudhistira91/technical-test-YC-W22/repository"
	"github.com/seanyudhistira91/technical-test-YC-W22/utils"
)

// Service interface defines business logic methods.
type Service interface {
	CreateUser(ctx context.Context, e entity.Users) error
	Login(ctx context.Context, e entity.Users) error
}

// concreteService implements the Service interface.
type concreteService struct {
	repo repository.Repository
}

// NewService creates a new instance of the concreteService.
func NewService(repo repository.Repository) Service {
	return &concreteService{repo: repo}
}

func (s *concreteService) CreateUser(ctx context.Context, e entity.Users) error {
	// Implement business logic here, potentially using the repository.
	err := s.repo.Create(ctx, e)
	if err != nil {
		return err
	}

	return nil
}

func (s *concreteService) Login(ctx context.Context, e entity.Users) error {
	// Implement business logic here, potentially using the repository.
	user, err := s.repo.GetByPhoneNumber(ctx, e.PhoneNumber)
	if err != nil {
		return err
	}

	// validate password
	err = utils.ComparePassword(user.HashPassword, e.HashPassword)
	if err != nil {
		return err
	}

	// TODO return JWT TOKEN
	return nil
}
