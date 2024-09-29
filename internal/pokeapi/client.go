package pokeapi

import (
	"net/http"
	"time"

	"github.com/Ammar4372/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, interval time.Duration) Client {
	cache := pokecache.NewCache(interval)
	return Client{
		httpClient: http.Client{
			Timeout: 3 * timeout,
		},
		cache: cache,
	}
}
