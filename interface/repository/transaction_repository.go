package repository

import (
	"github.com/gedearyarp/xendit-reconciliation/domain"
)

type TransactionRepository struct {
	csvHandler CSVHandler
}

func NewTransactionRepository(csvHandler CSVHandler) TransactionRepository {
	return TransactionRepository{
		csvHandler: csvHandler,
	}
}

func (repo TransactionRepository) ReadTransaction(fileName string) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	err := repo.csvHandler.ReadCSVFile(fileName, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
