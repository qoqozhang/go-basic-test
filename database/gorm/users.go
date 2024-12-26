package gorm

import (
	"context"
	"github.com/qoqozhang/go-basic-test.git/database/gorm/model"
)

func (db *database) CreateUser(ctx context.Context, user *model.User) error {
	tx := db.DB.WithContext(ctx)
	result := tx.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *database) SelectAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
