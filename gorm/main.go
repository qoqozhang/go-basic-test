package main

import (
	"context"
	"github.com/qoqozhang/go-basic-test.git/gorm/db"
	"github.com/qoqozhang/go-basic-test.git/gorm/router"
)

var (
	database db.Database
)

func init() {
	database = db.New(context.Background())
}

func main() {
	g := router.NewRouter(database)
	g.Run(":8080")

}
