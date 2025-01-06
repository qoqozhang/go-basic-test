package model

import (
	"time"
)

type Company struct {
	ID        uint       `gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Name      string     `json:"name,omitempty"`
}
