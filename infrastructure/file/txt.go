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
	sourceRecordText := fmt.Sprint("There are ", reconciliationSummary.SourceRecordProcessed, " source records processed.\n\n")
	discrepanciesText := ""
	
	countDiscrepancy := 0
	for discrepancy, totalDiscrepancy := range reconciliationSummary.MapDiscrepancies {
		countDiscrepancy += int(totalDiscrepancy)
		discrepanciesText += fmt.Sprint(totalDiscrepancy, " ", discrepancy, "\n")
	}

	discrepanciesText = fmt.Sprint("In total, there are ", reconciliationSummary.TotalReconciliation, " invalid proxies with ", countDiscrepancy, " discrepancies, concluded in a ", len(reconciliationSummary.MapDiscrepancies), " types of discrepancies: \n") + discrepanciesText

	fullText := dateRangeText + sourceRecordText + discrepanciesText
	fmt.Fprintf(file, fullText)

	return nil
}
