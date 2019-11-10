package cloudflare

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/himidori/cf-dynamic-dns/config"
	"github.com/himidori/cf-dynamic-dns/logs"
)

type CFRequester struct {
	httpClient *http.Client
	config     *config.Config
	workers    int
	workCh     chan *config.Site
	wg         *sync.WaitGroup
	logger     logs.Logger
}

type apiResult struct {
	Result json.RawMessage `json:"result"`
}

type dnsrecord struct {
	ID         string `json:"id"`
	RecordType string `json:"type"`
	DomainName string `json:"name"`
	Content    string `json:"content"`
	Proxiable  bool   `json:"proxiable"`
	Proxied    bool   `json:"proxied"`
}
