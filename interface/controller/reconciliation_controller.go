package controller

import (
	"os"

	"github.com/gedearyarp/xendit-recon-test/usecase"
)

type ReconciliationController struct {
	reconciliationInteractor usecase.ReconciliationInteractor
}

func NewReconciliationController(interactor usecase.ReconciliationInteractor) *ReconciliationController {
	return &ReconciliationController{interactor}
}

func (controller *ReconciliationController) ReconcilTransaction(proxyFileName string, sourceFileName string, reconciliationFileName string) error {
	if _, err := os.Stat(proxyFileName); err != nil {
		return err
	}
	if _, err := os.Stat(sourceFileName); err != nil {
		return err
	}

	err := controller.reconciliationInteractor.ReconcilTransaction(proxyFileName, sourceFileName, reconciliationFileName)
	if err != nil {
		return err
	}
	return nil
}
