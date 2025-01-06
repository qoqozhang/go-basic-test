package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qoqozhang/go-basic-test.git/gorm/api"
	"github.com/qoqozhang/go-basic-test.git/gorm/db"
)

func NewRouter(db db.Database) *gin.Engine {
	router := gin.Default()

	// user
	userAPI := api.UserAPI{
		DB: db,
	}
	user := router.Group("/user")
	user.GET("/:id", userAPI.Get)
	user.POST("", userAPI.Create)
	user.PUT("/:id", userAPI.Update)
	user.DELETE("/:id", userAPI.Delete)
	user.GET("/name/:username", userAPI.GetByUserName)
	router.GET("/users", userAPI.List)

	// company
	companyAPI := api.CompanyAPI{
		DB: db,
	}
	company := router.Group("/company")
	company.POST("", companyAPI.Create)
	router.GET("/companies", companyAPI.List)

	// creditCard
	creditAPI := api.CreditCardAPI{DB: db}
	creditCard := router.Group("/credit_card")
	creditCard.POST("", creditAPI.Create)
	router.GET("/credit_cards", creditAPI.List)

	return router
}
