package domain

type ReconciliationSummary struct {
	StartDate           string
	EndDate             string
	SourceDataProcessed int64
	TotalDiscrepancies  int64
}

type ReconciliationSummaryRepository interface {
	WriteSummaryReport(fileName string, summaryReport ReconciliationSummary) error
}
