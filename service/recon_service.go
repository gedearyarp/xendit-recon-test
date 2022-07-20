package service

import (
	m "recon_test/model"
)

type ReconService struct {
	Sources []m.Source
	Proxies []m.Proxy
}

type ReconData struct {
	SourceMap map[string]m.Source
	ProxyMap  map[string]m.Proxy
}

func (svc *ReconService) Perform() []m.ReconResult {
	reconData := ReconData{
		SourceMap: MapSources(svc.Sources),
		ProxyMap:  MapProxies(svc.Proxies),
	}

	return reconData.Compare()
}

func MapSources(sources []m.Source) map[string]m.Source {
	var sourceMap = make(map[string]m.Source)

	for _, source := range sources {
		sourceMap[source.ID] = source
	}

	return sourceMap
}

func MapProxies(proxies []m.Proxy) map[string]m.Proxy {
	var proxyMap = make(map[string]m.Proxy)

	for _, proxy := range proxies {
		proxyMap[proxy.ID] = proxy
	}

	return proxyMap
}

func (recon *ReconData) Compare() []m.ReconResult {
	var result []m.ReconResult

	proxyToSourceRecon := recon.ProxyToSource()
	result = append(result, proxyToSourceRecon...)

	sourceToProxyRecon := recon.SourceToProxy()
	result = append(result, sourceToProxyRecon...)

	return result
}

func (recon *ReconData) ProxyToSource() []m.ReconResult {
	var (
		result    []m.ReconResult
		reconData m.ReconResult
	)

	for key, value := range recon.ProxyMap {
		source, ok := recon.SourceMap[key]
		if !ok {
			reconData = m.ReconResult{
				ID:          value.ID,
				Amount:      value.Amount,
				Description: value.Description,
				Date:        value.Date,
				Remark:      m.SOURCE_NOT_FOUND,
			}
			result = append(result, reconData)
		} else {
			if value.Amount != source.Amount {
				reconData = m.ReconResult{
					ID:          value.ID,
					Amount:      value.Amount,
					Description: value.Description,
					Date:        value.Date,
					Remark:      m.AMOUNT_DIFF,
				}
				result = append(result, reconData)
			}

			if value.Date != source.Date {
				reconData = m.ReconResult{
					ID:          value.ID,
					Amount:      value.Amount,
					Description: value.Description,
					Date:        value.Date,
					Remark:      m.DATE_DIFF,
				}
				result = append(result, reconData)
			}

			if value.Description != source.Description {
				reconData = m.ReconResult{
					ID:          value.ID,
					Amount:      value.Amount,
					Description: value.Description,
					Date:        value.Date,
					Remark:      m.DATE_DIFF,
				}
				result = append(result, reconData)
			}
		}
	}
	return result
}

func (recon *ReconData) SourceToProxy() []m.ReconResult {
	var (
		result    []m.ReconResult
		reconData m.ReconResult
	)

	for key, value := range recon.SourceMap {
		_, ok := recon.ProxyMap[key]
		if !ok {
			reconData = m.ReconResult{
				ID:          value.ID,
				Amount:      value.Amount,
				Description: value.Description,
				Date:        value.Date,
				Remark:      m.PROXY_NOT_FOUND,
			}
			result = append(result, reconData)
		}
	}
	return result
}
