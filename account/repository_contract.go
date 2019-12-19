package account

import "service-account/models"

type RepositoryContract interface {
	GetAllAccounts() []models.Account
	GetAccount(id int64) models.Account
	CreateAccount(account models.Account) models.Account
	UpdateAccount(account models.Account, id int64) models.Account
	DeleteAccount(id int64) bool
}