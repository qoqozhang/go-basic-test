package db

import "github.com/qoqozhang/go-basic-test.git/gorm/model"

type CompanyDB interface {
	CreateCompany(*model.Company) error
	GetAllCompanies() ([]*model.Company, error)
	GetCompanyByID(int64) (*model.Company, error)
	UpdateCompany(*model.Company) error
	DeleteCompany(*model.Company) error
}

func (db *db) CreateCompany(company *model.Company) error {
	return db.DB.Create(company).Error
}
func (db *db) UpdateCompany(company *model.Company) error {
	return db.DB.Save(company).Error
}
func (db *db) DeleteCompany(company *model.Company) error {
	return db.DB.Delete(company).Error
}
func (db *db) GetAllCompanies() ([]*model.Company, error) {
	var companies []*model.Company
	err := db.DB.Find(&companies).Error
	return companies, err
}
func (db *db) GetCompanyByID(companyId int64) (*model.Company, error) {
	var company *model.Company
	err := db.DB.Where("id = ?", companyId).First(&company).Error
	return company, err
}
