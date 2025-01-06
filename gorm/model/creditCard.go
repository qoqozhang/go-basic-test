package model

import "time"

type CreditCard struct {
	ID        uint64    `gorm:"primaryKey;column:id"`
	CreatedAt time.Time `json:"-,omitempty"`
	UpdatedAt time.Time `json:"-,omitempty"`
	DeletedAt time.Time `gorm:"index" json:"-,omitempty"`
	Number    string    `gorm:"column:number" json:"number,omitempty"`
	UserID    uint64    `gorm:"column:user_id" json:"user_id,omitempty"`
}
