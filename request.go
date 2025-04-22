package http

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Http struct {
	host       string
	httpClient *http.Client
	header     map[string]string
}

func New(host string, header map[string]string) *Http {
	return &Http{
		host:       host,
		httpClient: &http.Client{Timeout: 5 * time.Minute},
		header:     header,
	}
}

func (h *Http) Host() string {
	return h.host
}

// Get http get 请求
func (h *Http) Get(ctx context.Context, param *url.Values) (*http.Response, error) {
	sign := "?"
	if strings.Contains(h.host, "?") {
		sign = "&"
	}
	if param != nil {
		h.host += sign + param.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, h.host, nil)
	if err != nil {
		return nil, err
	}

	if h.header != nil {
		for k, v := range h.header {
			req.Header.Add(k, v)
		}
	}

	return h.httpClient.Do(req)
}

// PostJson http post json 请求
func (h *Http) PostJson(ctx context.Context, param []byte) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.host, bytes.NewReader(param))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	if h.header != nil {
		for k, v := range h.header {
			req.Header.Add(k, v)
		}
	}

	return h.httpClient.Do(req)
}

// PostForm http post form 请求
func (h *Http) PostForm(ctx context.Context, param *url.Values) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.host, strings.NewReader(param.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if h.header != nil {
		for k, v := range h.header {
			req.Header.Add(k, v)
		}
	}

	return h.httpClient.Do(req)
}
