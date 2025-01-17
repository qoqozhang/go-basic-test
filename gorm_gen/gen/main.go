package main

import (
	"github.com/qoqozhang/go-basic-test.git/gorm_gen/db"
	"gorm.io/gen"
)

type Querier interface {
	// SELECT * from @@table where name = @@name and role = @@role
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}
type QuerierCompany interface {
	// select * from @@table where name = @@name
	FilterWithName(name string) ([]gen.M, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		//Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	gormdb := db.NewDB()
	g.UseDB(gormdb)
	//g.ApplyBasic(model.User{})
	userGenerator := g.GenerateModelAs("users", "User")
	g.ApplyBasic(userGenerator, g.GenerateModel("companies"))
	g.ApplyInterface(func(Querier) {}, userGenerator, g.GenerateModel("companies"))
	g.ApplyInterface(func(company QuerierCompany) {}, g.GenerateModel("companies"))
	g.Execute()
}
