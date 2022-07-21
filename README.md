# Coding Technical Assessment: Data Reconciliation

## How to use?
- Run main program  : `go run main.go` and check the `reconciliation.csv` at folder `./file`
- Run unit test     : `go test -v ./...`

## Folder structure
```.
├── README.md
├── file
│   ├── proxy.csv
│   ├── reconciliation.csv
│   └── source.csv
├── go.mod
├── go.sum
├── main.go
├── model
│   ├── recon_model.go
│   └── remark_model.go
├── service
│   ├── csv_service.go
│   └── recon_service.go
├── summary_report.txt
└── test
    ├── csv_service_test.go
    ├── file
    │   ├── proxy_test.csv
    │   ├── reconciliation_test.csv
    │   └── source_test.csv
    └── recon_service_test.go
```
