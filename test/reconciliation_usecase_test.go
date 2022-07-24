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

var (
	startDate = time.Date(2021, time.Month(7), 1, 0, 0, 0, 0, time.UTC)
	endDate   = time.Date(2021, time.Month(7), 31, 0, 0, 0, 0, time.UTC)
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

func TestReconciliationUsecase_AmountDifferent(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("file/test_amount_diff/proxy.csv", "file/test_amount_diff/source.csv", "file/test_amount_diff/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("file/test_amount_diff/reconciliation.csv", "file/test_amount_diff/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DateDifferent(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("file/test_date_diff/proxy.csv", "file/test_date_diff/source.csv", "file/test_date_diff/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("file/test_date_diff/reconciliation.csv", "file/test_date_diff/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DescrDifferent(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("file/test_descr_diff/proxy.csv", "file/test_descr_diff/source.csv", "file/test_descr_diff/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("file/test_descr_diff/reconciliation.csv", "file/test_descr_diff/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DateOutOfRange(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("file/test_date_out_of_range/proxy.csv", "file/test_date_out_of_range/source.csv", "file/test_date_out_of_range/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("file/test_date_out_of_range/reconciliation.csv", "file/test_date_out_of_range/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_SourceNotFound(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("file/test_source_not_found/proxy.csv", "file/test_source_not_found/source.csv", "file/test_source_not_found/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("file/test_source_not_found/reconciliation.csv", "file/test_source_not_found/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_Perfect(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("file/test_perfect/proxy.csv", "file/test_perfect/source.csv", "file/test_perfect/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("file/test_perfect/reconciliation.csv", "file/test_perfect/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DoubleRemark(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("file/test_double_remark/proxy.csv", "file/test_double_remark/source.csv", "file/test_double_remark/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("file/test_double_remark/reconciliation.csv", "file/test_double_remark/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_TenTransaction(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("file/test_double_remark/proxy.csv", "file/test_double_remark/source.csv", "file/test_double_remark/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("file/test_double_remark/reconciliation.csv", "file/test_double_remark/expected_reconciliation.csv"))
}