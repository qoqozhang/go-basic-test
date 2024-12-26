package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qoqozhang/go-basic-test.git/database/gorm"
	"github.com/qoqozhang/go-basic-test.git/database/router"
)

func main() {
	ctx := context.Background()
	//data := db.New(ctx)
	//defer data.Close()
	data := gorm.New(ctx)
	defer data.Close()

	g := gin.Default()
	router.NewRouter(g, data)
	g.Run(":8080")

}
