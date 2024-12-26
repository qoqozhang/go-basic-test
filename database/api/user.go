package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qoqozhang/go-basic-test.git/database/db"
	"github.com/qoqozhang/go-basic-test.git/database/gorm/model"
	"net/http"
)

type UserAPI struct {
	DB db.Database
}

type UserParameters struct {
	*model.User
}

func (api *UserAPI) Create(ctx *gin.Context) {
	// 获取数据
	// params :=gin.ShouldBindJson(UserParameters)
	params := &UserParameters{}
	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 写入数据库
	// api.DB.CreateUser(params)
	if err := params.SetPassword(params.User.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := api.DB.CreateUser(ctx, params.User); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 根据数据库的结果返回数据
	// gin.JSON(.....)
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully!", "code": 0})
}

func (api *UserAPI) List(ctx *gin.Context) {
	users, err := api.DB.SelectAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})

}
func (api *UserAPI) Update(ctx *gin.Context)    {}
func (api *UserAPI) Delete(ctx *gin.Context)    {}
func (api *UserAPI) DeleteAll(ctx *gin.Context) {}
func (api *UserAPI) Get(ctx *gin.Context)       {}
