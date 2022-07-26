package repository

import "github.com/gedearyarp/xendit-reconciliation/domain"

type CSVHandler interface {
	ReadCSVFile(fileName string, outData interface{}) error
	WriteCSVFile(fileName string, inData interface{}) error
}

type TXTHandler interface {
	WriteSummaryReport(fileName string, summaryReport domain.ReconciliationSummary) error
}