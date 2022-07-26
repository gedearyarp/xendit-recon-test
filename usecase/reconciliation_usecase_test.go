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
	file1, err := os.Open(fileName1)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file1)
	records1, _ := reader.ReadAll()

	file2, err := os.Open(fileName2)
	if err != nil {
		fmt.Println(err)
	}
	reader2 := csv.NewReader(file2)
	records2, _ := reader2.ReadAll()

	for i := range records1 {
		for j := range records1[i] {
			if records1[i][j] != records2[i][j] {
				return false
			}
		}
	}
	return true
}

func getReconciliationInteractor() ReconciliationInteractor {
	txtHandler := file.NewTXTHandler()
	csvHandler := file.NewCSVHandler()
	transactionRepo := repository.NewTransactionRepository(csvHandler)
	reconciliationRepo := repository.NewReconciliationRepository(csvHandler)
	reconciliationSummaryRepo := repository.NewReconciliationSummaryRepository(txtHandler)
	reconciliationInteractor := NewReconciliationInteractor(reconciliationRepo, transactionRepo, reconciliationSummaryRepo)
	return reconciliationInteractor
}

func TestReconciliationUsecase_AmountDifferent(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_amount_diff/proxy.csv", "../test/file/test_amount_diff/source.csv", "../test/file/test_amount_diff/reconciliation.csv", "../test/file/test_amount_diff/summary_report.txt", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_amount_diff/reconciliation.csv", "../test/file/test_amount_diff/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DateDifferent(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_date_diff/proxy.csv", "../test/file/test_date_diff/source.csv", "../test/file/test_date_diff/reconciliation.csv", "../test/file/test_date_diff/summary_report.txt", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_date_diff/reconciliation.csv", "../test/file/test_date_diff/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DescrDifferent(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_descr_diff/proxy.csv", "../test/file/test_descr_diff/source.csv", "../test/file/test_descr_diff/reconciliation.csv", "../test/file/test_descr_diff/summary_report.txt", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_descr_diff/reconciliation.csv", "../test/file/test_descr_diff/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DateOutOfRange(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_date_out_of_range/proxy.csv", "../test/file/test_date_out_of_range/source.csv", "../test/file/test_date_out_of_range/reconciliation.csv", "../test/file/test_date_out_of_range/summary_report.txt", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_date_out_of_range/reconciliation.csv", "../test/file/test_date_out_of_range/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_SourceNotFound(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_source_not_found/proxy.csv", "../test/file/test_source_not_found/source.csv", "../test/file/test_source_not_found/reconciliation.csv", "../test/file/test_source_not_found/summary_report.txt", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_source_not_found/reconciliation.csv", "../test/file/test_source_not_found/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_Perfect(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_perfect/proxy.csv", "../test/file/test_perfect/source.csv", "../test/file/test_perfect/reconciliation.csv", "../test/file/test_perfect/summary_report.txt", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_perfect/reconciliation.csv", "../test/file/test_perfect/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_DoubleRemark(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_double_remark/proxy.csv", "../test/file/test_double_remark/source.csv", "../test/file/test_double_remark/reconciliation.csv", "../test/file/test_double_remark/summary_report.txt", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_double_remark/reconciliation.csv", "../test/file/test_double_remark/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_TenTransaction(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_ten_transaction/proxy.csv", "../test/file/test_ten_transaction/source.csv", "../test/file/test_ten_transaction/reconciliation.csv", "../test/file/test_ten_transaction/summary_report.txt", startDate, endDate)
	assert.Nil(t, err)
	assert.True(t, compareCsv("../test/file/test_ten_transaction/reconciliation.csv", "../test/file/test_ten_transaction/expected_reconciliation.csv"))
}

func TestReconciliationUsecase_ErrorReadProxy(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/abcdefgh/proxy.csv", "../test/file/abcdefgh/source.csv", "../test/file/abcdefgh/reconciliation.csv", "../test/file/abcdefgh/summary_report.txt", startDate, endDate)
	assert.NotNil(t, err)
}

func TestReconciliationUsecase_ErrorReadSource(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_ten_transaction/proxy.csv", "../test/file/abcdefgh/source.csv", "../test/file/abcdefgh/reconciliation.csv", "../test/file/abcdefgh/summary_report.txt", startDate, endDate)
	assert.NotNil(t, err)
}

func TestReconciliationUsecase_ErrorWriteReconciliation(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_ten_transaction/proxy.csv", "../test/file/test_ten_transaction/source.csv", "../test/file/abcdefgh/reconciliation.csv", "../test/file/abcdefgh/summary_report.txt", startDate, endDate)
	assert.NotNil(t, err)
}

func TestReconciliationUsecase_ErrorWriteSummaryReport(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_ten_transaction/proxy.csv", "../test/file/test_ten_transaction/source.csv", "../test/file/test_ten_transaction/reconciliation.csv", "../test/file/abcdefgh/summary_report.txt", startDate, endDate)
	assert.NotNil(t, err)
}

func TestReconciliationUsecase_ErrorConvertDate(t *testing.T) {
	reconciliationInteractor := getReconciliationInteractor()

	err := reconciliationInteractor.ReconcileTransaction("../test/file/test_error/proxy.csv", "../test/file/test_error/source.csv", "../test/file/test_error/reconciliation.csv", "../test/file/test_error/summary_report.txt", startDate, endDate)
	assert.NotNil(t, err)
}
