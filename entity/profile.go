package entity

import "time"

// ResourceEntity represents the core business entity for the resource.
type Profiles struct {
	ID        int
	UserId    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
