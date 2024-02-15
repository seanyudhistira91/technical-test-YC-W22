package model

import "time"

// ResourceEntity represents the core business entity for the resource.
type Profiles struct {
	ID        int        `json:"id" gorm:"type:BIGINT;column:id;primaryKey"`
	UserId    int        `json:"userId" gorm:"type:BIGINT;column:user_id"`
	Name      string     `json:"name" gorm:"type:VARCHAR;column:name"`
	CreatedAt time.Time  `json:"createdAt" gorm:"type:timestamptz;column:created_at"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"type:timestamptz;column:updated_at"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"type:timestamptz;column:deleted_at"`
	User      Users      `json:"user" gorm:"foreignKey:UserId;references:ID"`
}
