package service

import (
	m "recon_test/model"
)

type ReconService struct {
	Sources []m.Source
	Proxies []m.Proxy
}

func (svc *ReconService) Compare() []m.ReconResult {
	var result []m.ReconResult

	return result
}
