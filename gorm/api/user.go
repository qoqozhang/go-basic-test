package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qoqozhang/go-basic-test.git/gorm/db"
	"github.com/qoqozhang/go-basic-test.git/gorm/model"
	"github.com/qoqozhang/go-basic-test.git/gorm/utils"
	"net/http"
	"strconv"
)

type UserAPI struct {
	DB db.Database
}
type userParams struct {
	Name             string `json:"name" binding:"required"`
	CreditCardNumber string `json:"credit_card" binding:"required"`
}

func (api *UserAPI) Create(c *gin.Context) {
	params := &userParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseJson(nil, err.Error(), http.StatusBadRequest))
		return
	}
	user := model.User{
		Name:       params.Name,
		CreditCard: model.CreditCard{Number: params.CreditCardNumber},
	}
	err := api.DB.CreateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseJson(nil, err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseJson(nil, nil, http.StatusCreated))
}
func (api *UserAPI) Delete(c *gin.Context) {
	var param string
	if param = c.Param("id"); param == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseJson(nil, "id is required", http.StatusBadRequest))
		return
	}
	id, _ := strconv.Atoi(param)
	if err := api.DB.DeleteUserById(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseJson(nil, err.Error(), http.StatusInternalServerError))
	} else {
		c.JSON(http.StatusOK, utils.ResponseJson(nil, nil, http.StatusOK))
	}
}
func (api *UserAPI) Update(c *gin.Context) {
	var err error
	paramId := c.Param("id")
	params := &userParams{}
	if err = c.ShouldBindJSON(params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseJson(nil, err.Error(), http.StatusBadRequest))
		return
	}
	var user *model.User
	id, _ := strconv.Atoi(paramId)
	if user, err = api.DB.GetUserById(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseJson(nil, err.Error(), http.StatusInternalServerError))
		return
	}
	user.Name = params.Name
	user.CreditCard.Number = params.CreditCardNumber
	if err = api.DB.UpdateUser(user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseJson(nil, err.Error(), http.StatusInternalServerError))
	} else {
		c.JSON(http.StatusOK, utils.ResponseJson(nil, nil, http.StatusOK))
	}
}
func (api *UserAPI) Get(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	user, err := api.DB.GetUserById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseJson(nil, err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseJson(user, nil, http.StatusOK))
}

func (api *UserAPI) GetByUserName(c *gin.Context) {
	paramUsername := c.Param("username")
	user, err := api.DB.GetUserByUsername(paramUsername)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseJson(nil, err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseJson(user, nil, http.StatusOK))
}
func (api *UserAPI) List(c *gin.Context) {
	users, err := api.DB.GetAllUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseJson(nil, err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseJson(users, nil, http.StatusOK))
}
