package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qoqozhang/go-basic-test.git/gorm/db"
	"github.com/qoqozhang/go-basic-test.git/gorm/model"
	"github.com/qoqozhang/go-basic-test.git/gorm/utils"
	"net/http"
)

type CreditCardAPI struct {
	DB db.Database
}
type creditCardParams struct {
	Number string `json:"number" binding:"required"`
}

func (api *CreditCardAPI) Create(c *gin.Context) {
	var params creditCardParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseJson(nil, err.Error(), http.StatusBadRequest))
		return
	}
	card := model.CreditCard{
		Number: params.Number,
	}
	if err := api.DB.CreateCreditCard(&card); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseJson(nil, err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseJson(nil, nil, http.StatusCreated))
}
func (api *CreditCardAPI) List(c *gin.Context) {
	cards, err := api.DB.ListCreditCards()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseJson(nil, err, http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseJson(cards, nil, http.StatusOK))
}
