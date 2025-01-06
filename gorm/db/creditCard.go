package db

import "github.com/qoqozhang/go-basic-test.git/gorm/model"

type creditCard interface {
	CreateCreditCard(*model.CreditCard) error
	ListCreditCards() ([]*model.CreditCard, error)
}

func (db *db) CreateCreditCard(creditCard *model.CreditCard) error {
	return db.DB.Create(creditCard).Error
}

func (db *db) ListCreditCards() ([]*model.CreditCard, error) {
	var cards []*model.CreditCard
	err := db.DB.Find(&cards).Error
	return cards, err
}
