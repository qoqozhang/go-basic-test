package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID         int        `json:"id" db:"id"`
	Username   string     `json:"username,omitempty" db:"username"`
	Email      string     `json:"email,omitempty" db:"email"`
	Password   string     `json:"password,-" db:"password"`
	LastLogin  *time.Time `json:"last_login,omitempty" db:"last_login"`
	FirstLogin *time.Time `json:"first_login,omitempty" db:"first_login"`
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	if u.Password != "" && len(u.Password) == 0 {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
