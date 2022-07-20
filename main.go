package main

import (
	"fmt"
	m "recon_test/model"
	svc "recon_test/service"
)

func main() {
	var (
		sourceCsv    svc.CsvService
		proxyCsv     svc.CsvService
		reconService svc.ReconService

		sources      []m.Source
		proxies      []m.Proxy
		reconResults []m.ReconResult
	)

	sourceCsv = svc.CsvService{FileName: "file/source.csv"}
	proxyCsv = svc.CsvService{FileName: "file/proxy.csv"}

	sources = sourceCsv.ReadSource()
	proxies = proxyCsv.ReadProxy()

	reconService = svc.ReconService{Sources: sources, Proxies: proxies}
	reconResults = reconService.Compare()

	fmt.Println(reconService)
}
