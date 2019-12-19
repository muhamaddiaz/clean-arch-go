package repository

import (
	"github.com/jinzhu/gorm"
	"service-account/account"
	"service-account/models"
)

type AccountRepository struct {
	*gorm.DB
}

func NewPostgresAccountRepository(db *gorm.DB) account.RepositoryContract {
	ar := &AccountRepository{db}

	db.AutoMigrate(models.Account{})

	return ar
}

func (ar *AccountRepository) GetAllAccounts() []models.Account {
	var accounts []models.Account

	ar.Find(&accounts)

	return accounts
}

func (ar *AccountRepository) GetAccount(id int64) models.Account {
	var accountInstance models.Account

	ar.Find(&accountInstance, id)

	return accountInstance
}

func (ar *AccountRepository) CreateAccount(account models.Account) models.Account {
	if !ar.NewRecord(account) {
		panic("Account telah terdaftar sebelumnya")
	}

	ar.Create(&account)

	return account
}

func (ar *AccountRepository) UpdateAccount(account models.Account, id int64) models.Account {
	ar.First(&account)

	ar.Save(&account)

	return account
}

func (ar *AccountRepository) DeleteAccount(id int64) bool {
	var accountInstance models.Account

	ar.Find(&accountInstance)

	ar.Delete(&accountInstance)

	return true
}