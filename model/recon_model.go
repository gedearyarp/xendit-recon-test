package model

type Proxy struct {
	ID          string `csv:"ID"`
	Amount      int64  `csv:"Amt"`
	Description string `csv:"Descr"`
	Date        string `csv:"Date"`
}

type Source struct {
	ID          string `csv:"ID"`
	Amount      int64  `csv:"Amount"`
	Description string `csv:"Description"`
	Date        string `csv:"Date"`
}

type ReconResult struct {
	ID          string `csv:"ID"`
	Amount      int64  `csv:"Amount"`
	Description string `csv:"Description"`
	Date        string `csv:"Date"`
	Remark      string `csv:"Remark"`
}
