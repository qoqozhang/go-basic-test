package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qoqozhang/go-basic-test.git/gorm/db"
	"github.com/qoqozhang/go-basic-test.git/gorm/model"
	"github.com/qoqozhang/go-basic-test.git/gorm/utils"
	"net/http"
)

type CompanyAPI struct {
	DB db.Database
}

type companyParams struct {
	Name string `json:"name"`
}

func (api *CompanyAPI) Create(c *gin.Context) {
	params := &companyParams{}
	if err := c.ShouldBind(params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseJson(nil, err.Error(), http.StatusBadRequest))
		return
	}
	company := model.Company{
		Name: params.Name,
	}
	if err := api.DB.CreateCompany(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseJson(nil, err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseJson(nil, nil, http.StatusCreated))
}
func (api *CompanyAPI) List(c *gin.Context) {
	companies, err := api.DB.GetAllCompanies()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseJson(nil, err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseJson(companies, nil, http.StatusOK))
}
