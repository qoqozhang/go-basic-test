package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qoqozhang/go-basic-test.git/database/api"
	"github.com/qoqozhang/go-basic-test.git/database/db"
)

func NewRouter(http *gin.Engine, database db.Database) {
	user := api.UserAPI{
		DB: database,
	}
	http.GET("/", index)
	http.GET("/users", user.List)
	http.POST("/users", user.Create)
	http.GET("/user/:id", user.Get)
}

func index(ctx *gin.Context) {
	ctx.String(200, "<h2>Pong</h2>")
}
