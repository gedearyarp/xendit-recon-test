package repository

import (
	"github.com/gedearyarp/xendit-reconciliation/domain"
)

type ReconciliationRepository struct {
	csvHandler CSVHandler
}

func NewReconciliationRepository(csvHandler CSVHandler) ReconciliationRepository {
	return ReconciliationRepository{
		csvHandler: csvHandler,
	}
}

func (repo ReconciliationRepository) WriteReconciliation(fileName string, reconciliations []domain.Reconciliation) error {
	return repo.csvHandler.WriteCSVFile(fileName, &reconciliations)
}
