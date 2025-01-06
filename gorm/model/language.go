package model

import "time"

type Language struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Name      string     `gorm:"unique;not null"`
}
