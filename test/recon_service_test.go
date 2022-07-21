package test

import (
	m "recon_test/model"
	svc "recon_test/service"
	"sort"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReconService_PerformSuccess(t *testing.T) {
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

	proxyCsv := svc.CsvService{FileName: "file/proxy_test.csv"}
	proxies, err := proxyCsv.ReadProxy()
	assert.Nil(t, err)

	sourceCsv := svc.CsvService{FileName: "file/source_test.csv"}
	sources, err := sourceCsv.ReadSource()
	assert.Nil(t, err)

	reconService := svc.ReconService{Sources: sources, Proxies: proxies}
	reconResult := reconService.Perform()
	assert.NotNil(t, reconResult)

	sort.Slice(recons, func(i int, j int) bool {
		return recons[i].ID < recons[j].ID
	})
	sort.Slice(reconResult, func(i int, j int) bool {
		return reconResult[i].ID < reconResult[j].ID
	})

	assert.Equal(t, reconResult, recons)
}

func TestReconService_AmountDifferent(t *testing.T) {
	proxies := []m.Proxy{
		{
			ID:          "aaaa",
			Amount:      10,
			Description: "A",
			Date:        "2021-06-30",
		},
	}
	sources := []m.Source{
		{
			ID:          "aaaa",
			Amount:      12,
			Description: "A",
			Date:        "2021-06-30",
		},
	}

	reconService := svc.ReconService{Sources: sources, Proxies: proxies}
	reconResult := reconService.Perform()

	assert.NotNil(t, reconResult)
	assert.Equal(t, reconResult[0].Remark, m.AMOUNT_DIFF)
}

func TestReconService_DescrDifferent(t *testing.T) {
	proxies := []m.Proxy{
		{
			ID:          "aaaa",
			Amount:      10,
			Description: "A",
			Date:        "2021-06-30",
		},
	}
	sources := []m.Source{
		{
			ID:          "aaaa",
			Amount:      10,
			Description: "B",
			Date:        "2021-06-30",
		},
	}

	reconService := svc.ReconService{Sources: sources, Proxies: proxies}
	reconResult := reconService.Perform()

	assert.NotNil(t, reconResult)
	assert.Equal(t, reconResult[0].Remark, m.DESCR_DIFF)
}

func TestReconService_DateDifferent(t *testing.T) {
	proxies := []m.Proxy{
		{
			ID:          "aaaa",
			Amount:      10,
			Description: "A",
			Date:        "2021-06-30",
		},
	}
	sources := []m.Source{
		{
			ID:          "aaaa",
			Amount:      10,
			Description: "A",
			Date:        "2021-06-29",
		},
	}

	reconService := svc.ReconService{Sources: sources, Proxies: proxies}
	reconResult := reconService.Perform()

	assert.NotNil(t, reconResult)
	assert.Equal(t, reconResult[0].Remark, m.DATE_DIFF)
}

func TestReconService_SourceNotFound(t *testing.T) {
	proxies := []m.Proxy{
		{
			ID:          "aaaa",
			Amount:      10,
			Description: "A",
			Date:        "2021-06-30",
		},
	}
	sources := []m.Source{}

	reconService := svc.ReconService{Sources: sources, Proxies: proxies}
	reconResult := reconService.Perform()

	assert.NotNil(t, reconResult)
	assert.Equal(t, reconResult[0].Remark, m.SOURCE_NOT_FOUND)
}

func TestReconService_ProxyNotFound(t *testing.T) {
	proxies := []m.Proxy{}
	sources := []m.Source{
		{
			ID:          "aaaa",
			Amount:      10,
			Description: "A",
			Date:        "2021-06-30",
		},
	}

	reconService := svc.ReconService{Sources: sources, Proxies: proxies}
	reconResult := reconService.Perform()

	assert.NotNil(t, reconResult)
	assert.Equal(t, reconResult[0].Remark, m.PROXY_NOT_FOUND)
}

func TestReconService_DoubleRemark(t *testing.T) {
	proxies := []m.Proxy{
		{
			ID:          "aaaa",
			Amount:      10,
			Description: "A",
			Date:        "2021-06-30",
		},
	}
	sources := []m.Source{
		{
			ID:          "aaaa",
			Amount:      11,
			Description: "B",
			Date:        "2021-06-30",
		},
	}

	reconService := svc.ReconService{Sources: sources, Proxies: proxies}
	reconResult := reconService.Perform()

	assert.NotNil(t, reconResult)
	assert.Equal(t, reconResult[0].Remark, m.AMOUNT_DIFF+m.DESCR_DIFF)
}

func TestReconService_PerfectReconciliation(t *testing.T) {
	proxies := []m.Proxy{
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
	}
	sources := []m.Source{
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
	}

	reconService := svc.ReconService{Sources: sources, Proxies: proxies}
	reconResult := reconService.Perform()

	assert.Nil(t, reconResult)
}
