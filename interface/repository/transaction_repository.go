package repository

import (
	"io/ioutil"

	"github.com/gedearyarp/xendit-reconciliation/domain"

	"github.com/gocarina/gocsv"
)

type TransactionRepository struct {
}

func NewTransactionRepository() TransactionRepository {
	return TransactionRepository{}
}

func (repo TransactionRepository) ReadTransaction(fileName string) ([]domain.Transaction, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var transactions []domain.Transaction
	_ = gocsv.UnmarshalBytes(bytes, &transactions)

	return transactions, nil
}
