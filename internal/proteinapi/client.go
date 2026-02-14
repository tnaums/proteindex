package proteinapi

import (
	"net/http"
	"time"

	"github.com/tnaums/proteindex/internal/proteincache"	
)

// Client -
type Client struct {
	cache      proteincache.Cache	
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: proteincache.NewCache(cacheInterval),		
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
