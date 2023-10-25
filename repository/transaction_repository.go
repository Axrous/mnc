package repository

import (
	"github.com/Axrous/mnc/helper"
	"github.com/Axrous/mnc/model/domain"
	"github.com/sonyarouje/simdb"
)

type TransactionRepository interface {
	Save(transaction domain.Transaction)
	GetByUserId(transaction domain.Transaction) ([]domain.Transaction, error)
}

type transactionRepositoryImpl struct {
	db *simdb.Driver
}

// GetByUserId implements TransactionRepository.
func (repo *transactionRepositoryImpl) GetByUserId(transaction domain.Transaction) ([]domain.Transaction, error) {
	// repo.db.Open(&domain.Merchant{}).Where("userId")
	panic("")
}

// Save implements TransactionRepository.
func (repo *transactionRepositoryImpl) Save(transaction domain.Transaction) {
	err := repo.db.Insert(&transaction)
	helper.PanicIfError(err)
}

func NewTransactionRepository(db *simdb.Driver) TransactionRepository {
	return &transactionRepositoryImpl{
		db: db,
	}
}
