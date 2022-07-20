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

	result = recon.ProxyToSource()
	result = append(result, recon.SourceToProxy()...)

	return result
}

func (recon *ReconData) ProxyToSource() []m.ReconResult {
	var result []m.ReconResult

	for key, value := range recon.ProxyMap {
		source, ok := recon.SourceMap[key]
		if !ok {
			result = append(result, RemarkProxy(value, m.SOURCE_NOT_FOUND))
		}
		if ok && value.Amount != source.Amount {
			result = append(result, RemarkProxy(value, m.AMOUNT_DIFF))
		}
		if ok && value.Date != source.Date {
			result = append(result, RemarkProxy(value, m.DATE_DIFF))
		}
		if ok && value.Description != source.Description {
			result = append(result, RemarkProxy(value, m.DESCR_DIFF))
		}
	}

	return result
}

func (recon *ReconData) SourceToProxy() []m.ReconResult {
	var result []m.ReconResult

	for key, value := range recon.SourceMap {
		_, ok := recon.ProxyMap[key]
		if !ok {
			result = append(result, RemarkSource(value, m.PROXY_NOT_FOUND))
		}
	}

	return result
}

func RemarkSource(source m.Source, remark string) m.ReconResult {
	return m.ReconResult{
		ID:          source.ID,
		Amount:      source.Amount,
		Description: source.Description,
		Date:        source.Date,
		Remark:      remark,
	}
}

func RemarkProxy(proxy m.Proxy, remark string) m.ReconResult {
	return m.ReconResult{
		ID:          proxy.ID,
		Amount:      proxy.Amount,
		Description: proxy.Description,
		Date:        proxy.Date,
		Remark:      remark,
	}
}
