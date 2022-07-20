package model

import (
	"time"
)

type Proxy struct {
	ID          string    `csv:"ID"`
	Amount      int64     `csv:"Amt"`
	Description string    `csv:"Descr"`
	Date        time.Time `csv:"Date"`
}

type Source struct {
	ID          string    `csv:"ID"`
	Amount      int64     `csv:"Amount"`
	Description string    `csv:"Description"`
	Date        time.Time `csv:"Date"`
}

type ReconResult struct {
	ID          string    `csv:"ID"`
	Amount      int64     `csv:"Amount"`
	Description string    `csv:"Description"`
	Date        time.Time `csv:"Date"`
	Remark      string    `csv:"Remarks"`
}
