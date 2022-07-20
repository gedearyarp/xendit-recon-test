package main

import (
	"fmt"
	m "recon_test/model"
	svc "recon_test/service"
)

func main() {
	var (
		sources []m.Source
		proxies []m.Proxy
	)

	sourceCsv := svc.CsvService{FileName: "file/source.csv"}
	proxyCsv := svc.CsvService{FileName: "file/proxy.csv"}

	sources = sourceCsv.ReadSource()
	proxies = proxyCsv.ReadProxy()

	fmt.Println(sources)
	fmt.Println(proxies)

}
