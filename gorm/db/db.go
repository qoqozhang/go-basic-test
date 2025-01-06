package db

import (
	"context"
	"github.com/qoqozhang/go-basic-test.git/gorm/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
)

type Database interface {
	io.Closer
	UserDB
	CompanyDB
	creditCard
}

type db struct {
	Ctx context.Context
	DB  *gorm.DB
}

func New(ctx context.Context) Database {
	sql, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	database := db{
		Ctx: ctx,
		DB:  sql,
	}
	if err := database.HealthCheck(); err != nil {
		log.Fatalf("database health check failed: %v", err)
	}
	sql.AutoMigrate(&model.User{}, &model.Language{})
	return &database
}

func (db *db) Close() error {
	sqlDb, _ := db.DB.DB()
	return sqlDb.Close()
}
func (db *db) HealthCheck() error {
	sqlDb, _ := db.DB.DB()
	return sqlDb.Ping()
}
