package repository

import (
	"context"

	"github.com/seanyudhistira91/technical-test-YC-W22/entity"
	"github.com/seanyudhistira91/technical-test-YC-W22/model"
	"gorm.io/gorm"
)

// concreteRepository implements the Repository interface.
type concreteRepository struct {
	// Include necessary dependencies here
	connection *gorm.DB
}

// NewRepository creates a new instance of the concreteRepository.
func NewRepository(db *gorm.DB) Repository {
	return &concreteRepository{
		connection: db,
	}
}

var _ Repository = &concreteRepository{}

func (r *concreteRepository) Create(ctx context.Context, e entity.Users) error {
	// Implement data retrieval logic here
	var user model.Users

	user.EntitieToModel(e)

	tx := r.connection.WithContext(ctx).Begin()
	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *concreteRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (e *entity.Users, err error) {
	// Implement data retrieval logic here
	var user model.Users

	q := r.connection.WithContext(ctx)
	q = q.Where("phone_number = ?", phoneNumber)
	err = q.Model(&user).Find(&user).Error
	if err != nil {
		return nil, err
	}

	entity := user.ModelToEntity()
	return &entity, nil
}
