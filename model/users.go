package model

import (
	"time"

	"github.com/seanyudhistira91/technical-test-YC-W22/entity"
)

// ResourceEntity represents the core business entity for the resource.
type Users struct {
	ID           int        `json:"id" gorm:"type:BIGINT;column:id;primaryKey"`
	PhoneNumber  string     `json:"phoneNumber" gorm:"type:VARCHAR;column:phone_number"`
	Email        string     `json:"email" gorm:"type:VARCHAR;column:email"`
	IsPremium    bool       `json:"isPremium" gorm:"type:BOOLEAN;column:is_premium"`
	HashPassword string     `json:"hashPassword" gorm:"type:VARCHAR;column:hash_password"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"type:timestamptz;column:created_at"`
	DeletedAt    *time.Time `json:"deletedAt" gorm:"type:timestamptz;column:deleted_at"`
	Profile      *Profiles  `json:"profile" gorm:"foreignKey:UserId"`
}

func (u *Users) EntitieToModel(e entity.Users) {
	profiles := Profiles{
		Name: e.Profile.Name,
	}
	u.PhoneNumber = e.PhoneNumber
	u.HashPassword = e.HashPassword
	u.Profile = &profiles
	u.Email = e.Email
}

func (u *Users) ModelToEntity() entity.Users {
	return entity.Users{
		ID:           u.ID,
		PhoneNumber:  u.PhoneNumber,
		Email:        u.Email,
		IsPremium:    u.IsPremium,
		HashPassword: u.HashPassword,
		CreatedAt:    u.CreatedAt,
	}

}
