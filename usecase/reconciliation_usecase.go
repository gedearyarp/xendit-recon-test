package usecase

import (
	"recon_test/domain"
)

type ReconciliationInteractor struct {
	reconciliationRepository domain.ReconciliationRepository
	transactionRepository    domain.TransactionRepository
}

func NewReconciliationInteractor(reconciliationRepository domain.ReconciliationRepository, transactionRepository domain.TransactionRepository) ReconciliationInteractor {
	return ReconciliationInteractor{
		reconciliationRepository: reconciliationRepository,
		transactionRepository:    transactionRepository,
	}
}

const (
	PROXY_NOT_FOUND  = "Data not found in Proxy (Only occured in Source);"
	SOURCE_NOT_FOUND = "Data not found in Source (Only occured in Proxy);"
	AMOUNT_DIFF      = "Amount between Proxy and Source is different;"
	DATE_DIFF        = "Date between Proxy and Source is different;"
	DESCR_DIFF       = "Description between Proxy and Source is different;"
)

func (interactor *ReconciliationInteractor) ReconcilTransaction(proxyFileName string, sourceFileName string, reconciliationFileName string) error {
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

	reconciliations := interactor.compare(mapProxies, mapSources)

	err = interactor.reconciliationRepository.WriteReconciliation(reconciliationFileName, reconciliations)
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

func (interactor *ReconciliationInteractor) compare(mapProxies map[string]domain.Transaction, mapSources map[string]domain.Transaction) []domain.Reconciliation {
	var result []domain.Reconciliation

	result = interactor.compareProxyToSource(mapProxies, mapSources)
	result = append(result, interactor.compareSourceToProxy(mapProxies, mapSources)...)

	return result
}

func (interactor *ReconciliationInteractor) compareProxyToSource(mapProxies map[string]domain.Transaction, mapSources map[string]domain.Transaction) []domain.Reconciliation {
	var result []domain.Reconciliation

	for id, proxy := range mapProxies {
		var remarks string
		source, ok := mapSources[id]
		if !ok {
			remarks = interactor.appendRemark(remarks, SOURCE_NOT_FOUND)
		}
		if ok && proxy.Amount != source.Amount {
			remarks = interactor.appendRemark(remarks, AMOUNT_DIFF)
		}
		if ok && proxy.Description != source.Description {
			remarks = interactor.appendRemark(remarks, DESCR_DIFF)
		}
		if ok && proxy.Date != source.Date {
			remarks = interactor.appendRemark(remarks, DATE_DIFF)
		}

		if remarks != "" {
			result = append(result, interactor.remarkTransaction(proxy, remarks))
		}
	}

	return result
}

func (interactor *ReconciliationInteractor) compareSourceToProxy(mapProxies map[string]domain.Transaction, mapSources map[string]domain.Transaction) []domain.Reconciliation {
	var result []domain.Reconciliation

	for id, source := range mapSources {
		_, ok := mapProxies[id]
		if !ok {
			result = append(result, interactor.remarkTransaction(source, PROXY_NOT_FOUND))
		}
	}

	return result
}

func (interactor *ReconciliationInteractor) remarkTransaction(source domain.Transaction, remark string) domain.Reconciliation {
	return domain.Reconciliation{
		ID:          source.ID,
		Amount:      source.Amount,
		Description: source.Description,
		Date:        source.Date,
		Remark:      remark,
	}
}

func (interactor *ReconciliationInteractor) appendRemark(remarks string, newRemark string) string {
	return remarks + newRemark
}
