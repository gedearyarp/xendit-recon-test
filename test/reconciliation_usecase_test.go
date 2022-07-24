package test

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/gedearyarp/xendit-reconciliation/infrastructure/file"
	"github.com/gedearyarp/xendit-reconciliation/interface/repository"
	"github.com/gedearyarp/xendit-reconciliation/usecase"

	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	startDate = "2021-07-01"
	endDate   = "2021-07-31"
)

func compareCsv(fileName1 string, fileName2 string) bool {
	file, err := os.Open(fileName1)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	file2, err := os.Open(fileName2)
	if err != nil {
		fmt.Println(err)
	}
	reader2 := csv.NewReader(file2)
	records2, _ := reader2.ReadAll()

	for i := range records {
		for j := range records[i] {
			if records[i][j] != records2[i][j] {
				return false
			}
		}
	}
	return true
}

func getReconciliationInteractor() usecase.ReconciliationInteractor {
	csvHandler := file.NewCSVHandler()
	transactionRepo := repository.NewTransactionRepository(csvHandler)
	reconciliationRepo := repository.NewReconciliationRepository(csvHandler)
	reconciliationInteractor := usecase.NewReconciliationInteractor(reconciliationRepo, transactionRepo)
	return reconciliationInteractor
}

func TestReconService_AmountDifferent(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()
	parsedStartDate, err := time.Parse("2006-01-02", startDate)
	assert.Nil(t, err)
	parsedEndDate, err := time.Parse("2006-01-02", endDate)
	assert.Nil(t, err)

	reconciledData := reconciliationInteractor.ReconcileTransaction("file/test_amount_different/proxy.csv", "file/test_amount_different/source.csv", "file/test_amount_different/reconciliation.csv", parsedStartDate, parsedEndDate)
	assert.NotNil(t, reconciledData)
	assert.True(t, compareCsv("file/test_amount_different/reconciliation.csv", "file/test_amount_different/expected_reconciliation.csv"))
}
