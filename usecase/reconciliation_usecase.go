package usecase

import (
	"fmt"
	"sort"
	t "time"

	"github.com/gedearyarp/xendit-reconciliation/domain"
)

type ReconciliationInteractor struct {
	reconciliationRepository        domain.ReconciliationRepository
	transactionRepository           domain.TransactionRepository
	reconciliationSummaryRepository domain.ReconciliationSummaryRepository
}

func NewReconciliationInteractor(reconciliationRepository domain.ReconciliationRepository, transactionRepository domain.TransactionRepository, reconciliationSummaryRepository domain.ReconciliationSummaryRepository) ReconciliationInteractor {
	return ReconciliationInteractor{
		reconciliationRepository:        reconciliationRepository,
		transactionRepository:           transactionRepository,
		reconciliationSummaryRepository: reconciliationSummaryRepository,
	}
}

const (
	SOURCE_NOT_FOUND  = "Transaction not found in Source (Only occured in Proxy);"
	AMOUNT_DIFF       = "Amount between Proxy and Source is different;"
	DATE_DIFF         = "Date between Proxy and Source is different;"
	DESCR_DIFF        = "Description between Proxy and Source is different;"
	DATE_OUT_OF_RANGE = "Transaction's date is out of range;"
)

func (interactor *ReconciliationInteractor) ReconcileTransaction(proxyFileName string, sourceFileName string, reconciliationFileName string, summaryReportFileName string, startDate t.Time, endDate t.Time) error {
	proxies, err := interactor.transactionRepository.ReadTransaction(proxyFileName)
	if err != nil {
		return err
	}
	sources, err := interactor.transactionRepository.ReadTransaction(sourceFileName)
	if err != nil {
		return err
	}

	mapProxies := interactor.mapTransactionById(proxies)
	mapSources := interactor.mapTransactionById(sources)

	reconciliations, totalDiscrepancies, err := interactor.compareTransaction(mapProxies, mapSources, startDate, endDate)
	if err != nil {
		return err
	}

	reconciliations = interactor.sortReconciliationById(reconciliations)
	reconciliationSummary := interactor.generateReconciliationSummary(startDate, endDate, int64(len(sources)), totalDiscrepancies)

	err = interactor.reconciliationRepository.WriteReconciliation(reconciliationFileName, reconciliations)
	if err != nil {
		return err
	}
	err = interactor.reconciliationSummaryRepository.WriteSummaryReport(summaryReportFileName, reconciliationSummary)
	if err != nil {
		return err
	}

	return nil
}

func (interactor *ReconciliationInteractor) mapTransactionById(transactions []domain.Transaction) map[string]domain.Transaction {
	var mapTransactions = make(map[string]domain.Transaction)

	for _, transaction := range transactions {
		mapTransactions[transaction.ID] = transaction
	}

	return mapTransactions
}

func (interactor *ReconciliationInteractor) compareTransaction(mapProxies map[string]domain.Transaction, mapSources map[string]domain.Transaction, startDate t.Time, endDate t.Time) ([]domain.Reconciliation, int64, error) {
	var (
		result             []domain.Reconciliation
		totalDiscrepancies int64
	)

	for id, proxy := range mapProxies {
		var remarks string

		remarks, err := interactor.remarkDateOutOfRange(remarks, proxy.Date, startDate, endDate)
		if err != nil {
			return nil, 0, err
		}

		source, ok := mapSources[id]
		if !ok {
			remarks = interactor.remarkSourceNotFound(remarks)
		} else {
			remarks = interactor.remarkDifferentField(remarks, proxy, source)
		}

		result = append(result, interactor.generateReconciliation(proxy, remarks))
		if remarks != "" {
			totalDiscrepancies++
		}
	}

	return result, totalDiscrepancies, nil
}

func (interactor *ReconciliationInteractor) remarkSourceNotFound(currentRemark string) string {
	return currentRemark + SOURCE_NOT_FOUND
}

func (interactor *ReconciliationInteractor) remarkDateOutOfRange(currentRemark string, date string, startDate t.Time, endDate t.Time) (string, error) {
	parsedDate, err := t.Parse("2006-01-02", date)
	if err != nil {
		return currentRemark, err
	}

	if parsedDate.After(endDate) || parsedDate.Before(startDate) {
		return currentRemark + DATE_OUT_OF_RANGE, nil
	}
	return currentRemark, nil
}

func (interactor *ReconciliationInteractor) remarkDifferentField(currentRemark string, proxy domain.Transaction, source domain.Transaction) string {
	if proxy.Amount != source.Amount {
		currentRemark = currentRemark + AMOUNT_DIFF
	}
	if proxy.Description != source.Description {
		currentRemark = currentRemark + DESCR_DIFF
	}
	if proxy.Date != source.Date {
		currentRemark = currentRemark + DATE_DIFF
	}
	return currentRemark
}

func (interactor *ReconciliationInteractor) generateReconciliation(source domain.Transaction, remark string) domain.Reconciliation {
	return domain.Reconciliation{
		ID:          source.ID,
		Amount:      source.Amount,
		Description: source.Description,
		Date:        source.Date,
		Remark:      remark,
	}
}

func (interactor *ReconciliationInteractor) generateReconciliationSummary(startDate t.Time, endDate t.Time, lenSource int64, totalDiscrepancies int64) domain.ReconciliationSummary {
	return domain.ReconciliationSummary{
		StartDate:           interactor.convertTimeToString(startDate),
		EndDate:             interactor.convertTimeToString(endDate),
		SourceDataProcessed: lenSource,
		TotalDiscrepancies:  totalDiscrepancies,
	}
}

func (interactor *ReconciliationInteractor) convertTimeToString(dateTime t.Time) string {
	return fmt.Sprint(dateTime.Day(), " ", dateTime.Month(), " ", dateTime.Year())
}

func (interactor *ReconciliationInteractor) sortReconciliationById(reconciliations []domain.Reconciliation) []domain.Reconciliation {
	sort.Slice(reconciliations, func(i int, j int) bool {
		return reconciliations[i].ID < reconciliations[j].ID
	})
	return reconciliations
}
