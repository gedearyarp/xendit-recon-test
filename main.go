package main

import (
	"log"

	"github.com/gedearyarp/xendit-reconciliation/interface/controller"
	"github.com/gedearyarp/xendit-reconciliation/interface/repository"
	"github.com/gedearyarp/xendit-reconciliation/usecase"
)

const (
	proxyFileName          = "file/transaction/proxy.csv"
	sourceFileName         = "file/transaction/source.csv"
	reconciliationFileName = "file/reconciliation/reconciliation.csv"
)

const (
	startDate = "2021-07-01"
	endDate   = "2021-07-31"
)

func getReconciliationController() controller.ReconciliationController {
	transactionRepo := repository.NewTransactionRepository()
	reconciliationRepo := repository.NewReconciliationRepository()
	reconciliationInteractor := usecase.NewReconciliationInteractor(reconciliationRepo, transactionRepo)
	reconciliationController := controller.NewReconciliationController(reconciliationInteractor)
	return *reconciliationController
}

func main() {
	reconciliationController := getReconciliationController()

	err := reconciliationController.ReconcileTransaction(proxyFileName, sourceFileName, reconciliationFileName, startDate, endDate)
	if err != nil {
		log.Fatal(err)
	}
}
