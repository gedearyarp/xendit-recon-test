package service

import (
	"io/ioutil"
	"os"
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

	return proxies
}

func (svc *CsvService) ReadSource() []m.Source {
	bytes, err := ioutil.ReadFile(svc.FileName)
	if err != nil {
		panic(err)
	}

	var sources []m.Source
	_ = gocsv.UnmarshalBytes(bytes, &sources)

	return sources
}

func (svc *CsvService) WriteResultRecon(reconResults []m.ReconResult) {
	file, _ := os.Create(svc.FileName)
	gocsv.MarshalFile(&reconResults, file)
}
