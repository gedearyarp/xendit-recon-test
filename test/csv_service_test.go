package test

import (
	m "github.com/gedearyarp/xendit-recon-test/model"
	svc "github.com/gedearyarp/xendit-recon-test/service"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsvService_ReadProxySuccess(t *testing.T) {
	proxy := []m.Proxy{
		{
			ID:          "aaaa",
			Amount:      10,
			Description: "A",
			Date:        "2021-06-30",
		},
		{
			ID:          "bbbb",
			Amount:      101,
			Description: "B",
			Date:        "2021-06-30",
		},
		{
			ID:          "cccc",
			Amount:      70,
			Description: "C",
			Date:        "2021-07-02",
		},
		{
			ID:          "dddd",
			Amount:      89,
			Description: "DD",
			Date:        "2021-08-03",
		},
		{
			ID:          "ffff",
			Amount:      24,
			Description: "F",
			Date:        "2021-09-03",
		},
	}
	proxyCsv := svc.CsvService{FileName: "file/proxy_test.csv"}
	result, err := proxyCsv.ReadProxy()

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, result, "Result must be not nil")
	assert.Equal(t, result, proxy, "Result must be equal to proxy data")
}

func TestCsvService_ReadProxyFail(t *testing.T) {
	proxyCsv := svc.CsvService{FileName: "file/proxy_test123.csv"}
	proxy, err := proxyCsv.ReadProxy()

	assert.NotNil(t, err, "proxy_test123.csv must be not found")
	assert.Nil(t, proxy, "Source must be nil")
}

func TestCsvService_ReadSourceSuccess(t *testing.T) {
	source := []m.Source{
		{
			ID:          "aaaa",
			Amount:      10,
			Description: "A",
			Date:        "2021-06-30",
		},
		{
			ID:          "bbbb",
			Amount:      100,
			Description: "B",
			Date:        "2021-06-30",
		},
		{
			ID:          "cccc",
			Amount:      70,
			Description: "C",
			Date:        "2021-07-01",
		},
		{
			ID:          "dddd",
			Amount:      89,
			Description: "D",
			Date:        "2021-08-03",
		},
		{
			ID:          "eeee",
			Amount:      71,
			Description: "E",
			Date:        "2021-09-03",
		},
	}

	sourceCsv := svc.CsvService{FileName: "file/source_test.csv"}
	result, err := sourceCsv.ReadSource()

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, result, "Result must be not nil")
	assert.Equal(t, result, source, "Result must be equal to proxy data")
}

func TestCsvService_ReadSourceFail(t *testing.T) {
	sourceCsv := svc.CsvService{FileName: "file/source_test123.csv"}
	source, err := sourceCsv.ReadProxy()

	assert.NotNil(t, err, "source_test123.csv must be not found")
	assert.Nil(t, source, "Source must be nil")
}

func TestCsvService_WriteResultReconSuccess(t *testing.T) {
	recons := []m.ReconResult{
		{
			ID:          "bbbb",
			Amount:      101,
			Description: "B",
			Date:        "2021-06-30",
			Remark:      m.AMOUNT_DIFF,
		},
		{
			ID:          "cccc",
			Amount:      70,
			Description: "C",
			Date:        "2021-07-02",
			Remark:      m.DATE_DIFF,
		},
		{
			ID:          "dddd",
			Amount:      89,
			Description: "DD",
			Date:        "2021-08-03",
			Remark:      m.DESCR_DIFF,
		},
		{
			ID:          "ffff",
			Amount:      24,
			Description: "F",
			Date:        "2021-09-03",
			Remark:      m.SOURCE_NOT_FOUND,
		},
		{
			ID:          "eeee",
			Amount:      71,
			Description: "E",
			Date:        "2021-09-03",
			Remark:      m.PROXY_NOT_FOUND,
		},
	}

	reconCsv := svc.CsvService{FileName: "file/reconciliation_test.csv"}
	_, err := reconCsv.WriteResultRecon(recons)
	assert.Nil(t, err)
}
