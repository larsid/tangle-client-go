package client

import (
	"net"
	"net/http"
	"sync"
	"time"

	iotago "github.com/iotaledger/iota.go/v2"
)

// DefaultHTTPTimeout is the per-request timeout applied to the shared HTTP client.
const DefaultHTTPTimeout = 30 * time.Second

var (
	httpOnce     sync.Once
	sharedClient *http.Client

	mu    sync.RWMutex
	cache = make(map[string]*iotago.NodeHTTPAPIClient)
)

func sharedHTTPClient() *http.Client {
	httpOnce.Do(func() {
		transport := &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 20,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		}

		sharedClient = &http.Client{
			Timeout: DefaultHTTPTimeout,
			Transport: &loggingTransport{
				base: transport,
			},
		}
	})
	return sharedClient
}

// ForURL returns a cached NodeHTTPAPIClient backed by the shared HTTP client pool.
func ForURL(nodeURL string) *iotago.NodeHTTPAPIClient {
	mu.RLock()
	if c, ok := cache[nodeURL]; ok {
		mu.RUnlock()
		return c
	}
	mu.RUnlock()

	mu.Lock()
	defer mu.Unlock()

	if c, ok := cache[nodeURL]; ok {
		return c
	}

	c := iotago.NewNodeHTTPAPIClient(
		nodeURL,
		iotago.WithNodeHTTPAPIClientHTTPClient(sharedHTTPClient()),
	)
	cache[nodeURL] = c
	return c
}
