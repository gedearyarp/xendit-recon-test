package usecase

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/gedearyarp/xendit-reconciliation/infrastructure/file"
	"github.com/gedearyarp/xendit-reconciliation/interface/repository"

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

func getReconciliationInteractor() ReconciliationInteractor {
	csvHandler := file.NewCSVHandler()
	transactionRepo := repository.NewTransactionRepository(csvHandler)
	reconciliationRepo := repository.NewReconciliationRepository(csvHandler)
	reconciliationInteractor := NewReconciliationInteractor(reconciliationRepo, transactionRepo)
	return reconciliationInteractor
}

func TestReconciliationUsecase_AmountDifferent(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_amount_diff/proxy.csv", "../test/file/test_amount_diff/source.csv", "../test/file/test_amount_diff/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_amount_diff/reconciliation.csv", "../test/file/test_amount_diff/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DateDifferent(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_date_diff/proxy.csv", "../test/file/test_date_diff/source.csv", "../test/file/test_date_diff/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_date_diff/reconciliation.csv", "../test/file/test_date_diff/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DescrDifferent(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_descr_diff/proxy.csv", "../test/file/test_descr_diff/source.csv", "../test/file/test_descr_diff/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_descr_diff/reconciliation.csv", "../test/file/test_descr_diff/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DateOutOfRange(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_date_out_of_range/proxy.csv", "../test/file/test_date_out_of_range/source.csv", "../test/file/test_date_out_of_range/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_date_out_of_range/reconciliation.csv", "../test/file/test_date_out_of_range/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_SourceNotFound(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_source_not_found/proxy.csv", "../test/file/test_source_not_found/source.csv", "../test/file/test_source_not_found/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_source_not_found/reconciliation.csv", "../test/file/test_source_not_found/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_Perfect(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_perfect/proxy.csv", "../test/file/test_perfect/source.csv", "../test/file/test_perfect/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_perfect/reconciliation.csv", "../test/file/test_perfect/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DoubleRemark(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_double_remark/proxy.csv", "../test/file/test_double_remark/source.csv", "../test/file/test_double_remark/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_double_remark/reconciliation.csv", "../test/file/test_double_remark/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_TenTransaction(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_ten_transaction/proxy.csv", "../test/file/test_ten_transaction/source.csv", "../test/file/test_ten_transaction/reconciliation.csv", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_ten_transaction/reconciliation.csv", "../test/file/test_ten_transaction/expected_reconciliation.csv"))
}
