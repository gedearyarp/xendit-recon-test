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

func (svc *CsvService) ReadProxy() ([]m.Proxy, error) {
	bytes, err := ioutil.ReadFile(svc.FileName)
	if err != nil {
		return nil, err
	}

	var proxies []m.Proxy
	_ = gocsv.UnmarshalBytes(bytes, &proxies)

	return proxies, nil
}

func (svc *CsvService) ReadSource() ([]m.Source, error) {
	bytes, err := ioutil.ReadFile(svc.FileName)
	if err != nil {
		return nil, err
	}

	var sources []m.Source
	_ = gocsv.UnmarshalBytes(bytes, &sources)

	return sources, nil
}

func (svc *CsvService) WriteResultRecon(reconResults []m.ReconResult) (bool, error) {
	file, err := os.Create(svc.FileName)
	if err != nil {
		return false, err
	}
	gocsv.MarshalFile(&reconResults, file)
	return true, nil
}
