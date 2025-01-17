package db

import (
	"github.com/qoqozhang/go-basic-test.git/gorm/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func NewDB() *gorm.DB {
	sql, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db, _ := sql.DB()
	if err := db.Ping(); err != nil {
		log.Fatalf("db.Ping() error(%v)", err)
	}
	sql.AutoMigrate(model.User{}, model.Company{})
	return sql
}
