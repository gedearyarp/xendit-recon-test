package repository

import (
	"os"

	"github.com/gedearyarp/xendit-recon-test/domain"

	"github.com/gocarina/gocsv"
)

type ReconciliationRepository struct {
}

func NewReconciliationRepository() ReconciliationRepository {
	return ReconciliationRepository{}
}

func (repo ReconciliationRepository) WriteReconciliation(fileName string, reconciliations []domain.Reconciliation) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	gocsv.MarshalFile(&reconciliations, file)
	return nil
}
