package theoneapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	apiKey     string
}

func NewClient(timeout time.Duration, apiKey string) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		apiKey: apiKey,
	}
}
