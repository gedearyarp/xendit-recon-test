package main

import (
	"fmt"

	svc "recon_test/service"
)

func main() {
	cc := svc.CsvService{FileName: "file/proxy.csv"}
	ress := cc.ReadProxy()
	fmt.Println(ress)
}
