package file

import (
	"fmt"
	"os"

	"github.com/gedearyarp/xendit-reconciliation/domain"
)

type txtHandler struct{}

func NewTXTHandler() *txtHandler {
	return &txtHandler{}
}

func (h *txtHandler) WriteSummaryReport(fileName string, reconciliationSummary domain.ReconciliationSummary) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	dateRangeText := fmt.Sprint("Reconciliation Summary Report (", reconciliationSummary.StartDate, " - ", reconciliationSummary.EndDate, ")\n\n")
	sourceRecordText := fmt.Sprint("There is ", reconciliationSummary.SourceDataProcessed, " source records processed.\n")
	discrepanciesText := fmt.Sprint("In total ", reconciliationSummary.TotalDiscrepancies, " discrepancies with 5 different type of discrepancy: ")

	fullText := dateRangeText + sourceRecordText + discrepanciesText
	fmt.Fprintf(file, fullText)

	return nil
}
