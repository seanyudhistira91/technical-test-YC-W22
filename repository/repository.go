// repository/repository.go
package repository

import (
	"context"

	"github.com/seanyudhistira91/technical-test-YC-W22/entity"
)

// Repository interface defines methods to interact with data storage.
type Repository interface {
	// Users
	Create(ctx context.Context, e entity.Users) error
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.Users, error)
}
