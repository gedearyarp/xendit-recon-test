package main

import (
	"fmt"
	m "recon_test/model"
	svc "recon_test/service"
)

func main() {
	var (
		sources      []m.Source
		proxies      []m.Proxy
		reconResults []m.ReconResult
	)

	sourceCsv := svc.CsvService{FileName: "file/source.csv"}
	proxyCsv := svc.CsvService{FileName: "file/proxy.csv"}

	sources, err := sourceCsv.ReadSource()
	if err != nil {
		panic(err)
	}
	proxies, err = proxyCsv.ReadProxy()
	if err != nil {
		panic(err)
	}

	reconService := svc.ReconService{Sources: sources, Proxies: proxies}
	reconResults = reconService.Perform()

	reconResultCsv := svc.CsvService{FileName: "file/reconciliation.csv"}
	_, err = reconResultCsv.WriteResultRecon(reconResults)
	if err != nil {
		panic(err)
	}

	fmt.Println("Reconciliation Success, please check at ./file/reconciliation")
}
