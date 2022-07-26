package controller

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/gedearyarp/xendit-reconciliation/usecase"
)

type ReconciliationController struct {
	reconciliationInteractor usecase.ReconciliationInteractor
}

func NewReconciliationController(interactor usecase.ReconciliationInteractor) *ReconciliationController {
	return &ReconciliationController{interactor}
}

const (
	INVALID_FILE_EXTENSION = "Invalid file extension"
)

func (controller *ReconciliationController) ReconcileTransaction(proxyFileName string, sourceFileName string, reconciliationFileName string, summaryReportFileName string, startDate string, endDate string) error {
	if _, err := os.Stat(proxyFileName); err != nil {
		return err
	}
	if _, err := os.Stat(sourceFileName); err != nil {
		return err
	}

	err := controller.validateFileExtension(proxyFileName, sourceFileName, reconciliationFileName, summaryReportFileName)
	if err != nil {
		return err
	}

	parsedStartDate, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return err
	}
	parsedEndDate, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return err
	}

	err = controller.reconciliationInteractor.ReconcileTransaction(proxyFileName, sourceFileName, reconciliationFileName, summaryReportFileName, parsedStartDate, parsedEndDate)
	if err != nil {
		return err
	}

	return nil
}

func (controller *ReconciliationController) validateFileExtension(proxyFileName string, sourceFileName string, reconciliationFileName string, summaryReportFileName string) error {
	if filepath.Ext(proxyFileName) != ".csv" ||
		filepath.Ext(sourceFileName) != ".csv" ||
		filepath.Ext(reconciliationFileName) != ".csv" ||
		filepath.Ext(summaryReportFileName) != ".txt" {
		return errors.New(INVALID_FILE_EXTENSION)
	}
	return nil
}
