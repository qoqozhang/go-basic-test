package db

import (
	"context"
	"github.com/qoqozhang/go-basic-test.git/database/gorm/model"
)

type UsersDB interface {
	CreateUser(ctx context.Context, user *model.User) error
	SelectAllUsers(ctx context.Context) ([]model.User, error)
}

var createUserQuery = `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

func (db *database) CreateUser(ctx context.Context, user *model.User) error {
	_, err := db.DB.Exec(createUserQuery, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (db *database) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}

func (db *database) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}

func (db *database) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (*model.User, error) {
	return nil, nil
}

func (db *database) UpdateUser(ctx context.Context, user *model.User) error {
	return nil
}
func (db *database) DeleteUser(ctx context.Context, user *model.User) error {
	return nil
}

func (db *database) DeleteUserByUsername(ctx context.Context, username string) error {
	return nil
}
func (db *database) DeleteUserByEmail(ctx context.Context, email string) error {
	return nil
}
func (db *database) SelectUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}
func (db *database) SelectUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}
func (db *database) SelectUserByEmailAndPassword(ctx context.Context, email string, password string) (*model.User, error) {
	return nil, nil
}
func (db *database) SelectAllUsers(ctx context.Context) ([]model.User, error) {
	/*
		var users []model.User
		rows, err := db.DB.QueryContext(ctx, "SELECT * FROM users")
		if err != nil {
			return users, err
		}
		defer rows.Close()
		for rows.Next() {
			user := model.User{}
			err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.LastLogin, &user.FirstLogin)
			if err != nil {
				return users, err
			}
			users = append(users, user)
		}

	*/
	return nil, nil
}
