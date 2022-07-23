package domain

type Transaction struct {
	ID          string `csv:"ID"`
	Amount      int64  `csv:"Amt,Amount"`
	Description string `csv:"Descr,Description"`
	Date        string `csv:"Date"`
}

type TransactionRepository interface {
	ReadTransaction(fileName string) ([]Transaction, error)
}
