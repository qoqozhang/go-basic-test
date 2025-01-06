package db

import (
	"github.com/qoqozhang/go-basic-test.git/gorm/model"
	"gorm.io/gorm"
)

type UserDB interface {
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	DeleteUserById(int) error
	GetAllUsers() ([]*model.User, error)
	GetUserByUsername(string) (*model.User, error)
	GetUserById(int) (*model.User, error)
}

// CreateUser 创建user 和关联的creditCard信息
func (db *db) CreateUser(user *model.User) error {
	return db.DB.Create(user).Error
}

// UpdateUser 更新用户
func (db *db) UpdateUser(user *model.User) error {
	return db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(user).Error
}

// DeleteUserById 根据userId删除用户已经关联的信用卡信息
func (db *db) DeleteUserById(id int) error {
	return db.DB.Model(&model.User{}).Where("id = ?", id).Delete(&model.User{ID: uint(id)}).Error
}

func (db *db) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	err := db.DB.Model(&model.User{}).Preload("CreditCard").Find(&users).Error
	return users, err
}
func (db *db) GetUserByUsername(username string) (*model.User, error) {
	var user *model.User
	err := db.DB.Model(&model.User{}).Preload("CreditCard").Where("name = ?", username).Find(&user).Error
	return user, err
}
func (db *db) GetUserById(id int) (*model.User, error) {
	var user *model.User
	err := db.DB.Model(&model.User{}).Preload("CreditCard").First(&user, id).Error
	return user, err
}
