package entity

import "time"

// ResourceEntity represents the core business entity for the resource.
type Users struct {
	ID           int
	PhoneNumber  string
	Email        string
	IsPremium    bool
	HashPassword string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	Profile      *Profiles
}
