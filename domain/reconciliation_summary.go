package domain

type ReconciliationSummary struct {
	StartDate             string
	EndDate               string
	SourceRecordProcessed int64
	TotalReconciliation   int64
	MapDiscrepancies      map[string]int64
}

type ReconciliationSummaryRepository interface {
	WriteSummaryReport(fileName string, summaryReport ReconciliationSummary) error
}
