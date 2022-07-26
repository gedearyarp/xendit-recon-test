package main

import (
	"fmt"
	"log"

	"github.com/gedearyarp/xendit-reconciliation/infrastructure/file"
	"github.com/gedearyarp/xendit-reconciliation/interface/controller"
	"github.com/gedearyarp/xendit-reconciliation/interface/repository"
	"github.com/gedearyarp/xendit-reconciliation/usecase"
)

const (
	proxyFileName          = "file/transaction/proxy.csv"
	sourceFileName         = "file/transaction/source.csv"
	reconciliationFileName = "file/reconciliation/reconciliation.csv"
	summaryReportFileName  = "file/summary_report.txt"
)

const (
	startDate = "2021-07-01"
	endDate   = "2021-07-31"
)

func getReconciliationController() controller.ReconciliationController {
	txtHandler := file.NewTXTHandler()
	csvHandler := file.NewCSVHandler()
	transactionRepo := repository.NewTransactionRepository(csvHandler)
	reconciliationRepo := repository.NewReconciliationRepository(csvHandler)
	reconciliationSummaryRepo := repository.NewReconciliationSummaryRepository(txtHandler)
	reconciliationInteractor := usecase.NewReconciliationInteractor(reconciliationRepo, transactionRepo, reconciliationSummaryRepo)
	reconciliationController := controller.NewReconciliationController(reconciliationInteractor)
	return *reconciliationController
}

func readInputFileName() (string, string, string, string) {
	var (
		proxyFileName          string
		sourceFileName         string
		reconciliationFileName string
		summaryReportFileName  string
	)

	fmt.Print("Enter Proxy's file path: ")
	fmt.Scanln(&proxyFileName)

	fmt.Print("Enter Source's file path: ")
	fmt.Scanln(&sourceFileName)

	fmt.Print("Enter file path for the reconciliation result (.csv): ")
	fmt.Scanln(&reconciliationFileName)

	fmt.Print("Enter file path for the reconciliation summary report (.txt): ")
	fmt.Scanln(&summaryReportFileName)
	return proxyFileName, sourceFileName, reconciliationFileName, summaryReportFileName
}

func readInputRangeDate() (string, string) {
	var (
		startDate string
		endDate   string
	)

	fmt.Print("Enter start date (ex. 2021-07-01): ")
	fmt.Scanln(&startDate)

	fmt.Print("Enter end date (ex. 2021-07-31): ")
	fmt.Scanln(&endDate)
	return startDate, endDate
}

func main() {
	reconciliationController := getReconciliationController()

	// proxyFileName, sourceFileName, reconciliationFileName, summaryReportFileName := readInputFileName()
	// startDate, endDate := readInputRangeDate()

	err := reconciliationController.ReconcileTransaction(proxyFileName, sourceFileName, reconciliationFileName, summaryReportFileName, startDate, endDate)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Reconciliation Success!")
	}
}
