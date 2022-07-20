package service

import (
	"fmt"
	"io/ioutil"
	m "recon_test/model"

	"github.com/gocarina/gocsv"
)

type CsvService struct {
	FileName string
}

func (svc *CsvService) ReadProxy() []m.Proxy {
	bytes, err := ioutil.ReadFile(svc.FileName)
	if err != nil {
		panic(err)
	}

	var proxies []m.Proxy
	_ = gocsv.UnmarshalBytes(bytes, &proxies)

	fmt.Println(svc.FileName)
	
	return proxies
}