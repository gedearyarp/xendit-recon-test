package test

import (
	m "recon_test/model"
	svc "recon_test/service"

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
			Amount:      100,
			Description: "B",
			Date:        "2021-06-30",
		},
		{
			ID:          "cccc",
			Amount:      1000,
			Description: "C",
			Date:        "2021-07-01",
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
	_, err := proxyCsv.ReadProxy()

	assert.NotNil(t, err, "proxy_test123.csv must be not found")
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
			Amount:      1000,
			Description: "C",
			Date:        "2021-07-01",
		},
		{
			ID:          "cccc",
			Amount:      1000,
			Description: "C",
			Date:        "2021-07-01",
		},
		{
			ID:          "cccc",
			Amount:      1000,
			Description: "C",
			Date:        "2021-07-01",
		},
	}

}
