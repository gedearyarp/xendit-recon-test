package repository

import "github.com/gedearyarp/xendit-reconciliation/domain"

type ReconciliationSummaryRepository struct {
	txtHandler TXTHandler
}

func NewReconciliationSummaryRepository(txtHandler TXTHandler) ReconciliationSummaryRepository {
	return ReconciliationSummaryRepository{
		txtHandler: txtHandler,
	}
}

func (repo ReconciliationSummaryRepository) WriteSummaryReport(fileName string, summaryReport domain.ReconciliationSummary) error {
	err := repo.txtHandler.WriteSummaryReport(fileName, summaryReport)
	if err != nil {
		return err
	}

	return nil
}
