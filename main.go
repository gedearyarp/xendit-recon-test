package main

import (
	"log"

	"recon_test/interface/controller"
	"recon_test/interface/repository"
	"recon_test/usecase"
)

const (
	proxyFileName          = "file/proxy.csv"
	sourceFileName         = "file/source.csv"
	reconciliationFileName = "file/reconciliation.csv"
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

	err := reconciliationController.ReconcilTransaction(proxyFileName, sourceFileName, reconciliationFileName)
	if err != nil {
		log.Fatal(err)
	}
}
