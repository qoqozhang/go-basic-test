package db

import (
	"github.com/qoqozhang/go-basic-test.git/gorm/model"
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

// UpdateUser 更新用户, 此操作会更新user信息，但是由于是一对多的原因，不会删除原来的credit_card信息，而是增加两条,所以采用先更新user信息，然后替换user的关联方式
func (db *db) UpdateUser(user *model.User) error {
	err := db.DB.Save(user).Error
	if err != nil {
		return err
	}
	err = db.DB.Model(user).Association("Languages").Replace(user.Languages)
	return err
}

// DeleteUserById 根据userId删除用户已经关联的信用卡信息
func (db *db) DeleteUserById(id int) error {
	return db.DB.Select("Languages").Delete(&model.User{ID: uint(id)}).Error
}

func (db *db) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	err := db.DB.Model(&model.User{}).Preload("Languages").Find(&users).Error
	return users, err
}
func (db *db) GetUserByUsername(username string) (*model.User, error) {
	var user *model.User
	err := db.DB.Model(&model.User{}).Preload("Languages").Where("name = ?", username).Find(&user).Error
	return user, err
}
func (db *db) GetUserById(id int) (*model.User, error) {
	var user *model.User
	err := db.DB.Model(&model.User{}).Preload("Languages").First(&user, id).Error
	return user, err
}
