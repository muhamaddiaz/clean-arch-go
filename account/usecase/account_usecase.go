package usecase

import (
	"service-account/account"
	"service-account/models"
)

type AccountUsecase struct {
	ar account.RepositoryContract
}

func NewAccountUsecase(acRep account.RepositoryContract) AccountUsecase {
	aUsecase := AccountUsecase{ar:acRep}

	return aUsecase
}

func (ac *AccountUsecase) GetAllAccounts() []models.Account {

	return ac.ar.GetAllAccounts()
}

func (ac *AccountUsecase) GetAccount(id int64) models.Account {

	return ac.ar.GetAccount(id)
}

func (ac *AccountUsecase) CreateAccount(account models.Account) models.Account {

	accountInstance := ac.ar.CreateAccount(account)

	return accountInstance
}

func (ac *AccountUsecase) UpdateAccount(account models.Account, id int64) models.Account {
	return account
}

func (ac *AccountUsecase) DeleteAccount(id int64) bool {
	return true
}
