package domain

type Reconciliation struct {
	ID          string `csv:"ID"`
	Amount      int64  `csv:"Amount"`
	Description string `csv:"Description"`
	Date        string `csv:"Date"`
	Remark      string `csv:"Remark"`
}

type ReconciliationRepository interface {
	WriteReconciliation(fileName string, reconciliation []Reconciliation) error
}
