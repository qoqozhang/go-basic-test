package gorm

import (
	"context"
	"github.com/qoqozhang/go-basic-test.git/database/db"
	"github.com/qoqozhang/go-basic-test.git/database/gorm/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type database struct {
	Ctx context.Context
	DB  *gorm.DB
}

func New(ctx context.Context) db.Database {
	db, err := gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	db.AutoMigrate(&model.User{})
	return &database{Ctx: ctx, DB: db}
}
func (db *database) Close() error {
	return db.Close()
}
