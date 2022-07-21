# Coding Technical Assessment: Data Reconciliation

## How to use?
- Run main program  : `go run main.go` and check the `reconciliation.csv` at folder `./file`
- Run unit test     : `go test -v ./...`

## Demo
### Run main program
![recon run demo](https://user-images.githubusercontent.com/71829426/180240425-bd2a3c8a-88bc-4b50-8abd-52a78c91db29.gif)
### Run unit test
![recon test demo](https://user-images.githubusercontent.com/71829426/180241701-abcd0e4d-215a-4b1e-8b72-fdf97346eceb.gif)


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
