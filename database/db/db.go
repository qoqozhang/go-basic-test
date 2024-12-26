package db

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"time"
)

type Database interface {
	io.Closer
	UsersDB
}

type database struct {
	Ctx context.Context
	DB  *sql.DB
}

func New(ctx context.Context) Database {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	_database := database{
		Ctx: ctx,
		DB:  db,
	}

	// 检查数据库是否可以成功访问
	_database.Ping()

	// migrate 创建数据库表格
	// TODO
	migrateDb(db)
	return &_database
}

func (db *database) Ping() {
	var ready = make(chan struct{})
	go func(ready chan struct{}) {
		err := db.DB.PingContext(db.Ctx)
		if err != nil {
			log.Fatalf("Error pinging database: %v", err)
		}
		close(ready)
	}(ready)
	select {
	case <-ready:
		log.Println("Database ready.")
	case <-time.After(10 * time.Second):
		log.Fatalf("Error pinging database: timeout")
	}
}

func (db *database) Close() error {
	return db.DB.Close()
}
