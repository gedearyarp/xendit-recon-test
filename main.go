package main

import (
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

	sources = sourceCsv.ReadSource()
	proxies = proxyCsv.ReadProxy()

	reconService := svc.ReconService{Sources: sources, Proxies: proxies}
	reconResults = reconService.Perform()

	reconResultCsv := svc.CsvService{FileName: "reconciliation.csv"}
	reconResultCsv.WriteResultRecon(reconResults)
}
