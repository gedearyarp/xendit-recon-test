package controller

import (
	"os"
	"time"

	"github.com/gedearyarp/xendit-reconciliation/usecase"
)

type ReconciliationController struct {
	reconciliationInteractor usecase.ReconciliationInteractor
}

func NewReconciliationController(interactor usecase.ReconciliationInteractor) *ReconciliationController {
	return &ReconciliationController{interactor}
}

func (controller *ReconciliationController) ReconcileTransaction(proxyFileName string, sourceFileName string, reconciliationFileName string, startDate string, endDate string) error {
	if _, err := os.Stat(proxyFileName); err != nil {
		return err
	}
	if _, err := os.Stat(sourceFileName); err != nil {
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

	err = controller.reconciliationInteractor.ReconcileTransaction(proxyFileName, sourceFileName, reconciliationFileName, parsedStartDate, parsedEndDate)
	if err != nil {
		return err
	}

	return nil
}