package cloudflare

import (
	"net/http"
	"sync"
	"time"

	"github.com/himidori/cf-dynamic-dns/config"
	"github.com/himidori/cf-dynamic-dns/logs"
)

func NewCFRequester(cfg *config.Config, workers int) *CFRequester {
	return &CFRequester{
		httpClient: &http.Client{Timeout: time.Second * 10},
		config:     cfg,
		workers:    workers,
		workCh:     make(chan config.Site),
		wg:         &sync.WaitGroup{},
		logger:     logs.NewLogger(),
	}
}
