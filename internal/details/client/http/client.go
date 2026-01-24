package http

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type Client struct {
	client *http.Client
}

func NewClient(cfg client.Config) *Client {
	// TODO: cfg
	return &Client{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: func(r *http.Request) (*url.URL, error) {
					if cfg.ProxyURL() == "" {
						return nil, nil
					}

					return url.Parse(cfg.ProxyURL())
				},
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					conn, err := (&net.Dialer{
						Timeout:   3 * time.Second,
						KeepAlive: 3 * time.Second,
					}).DialContext(ctx, network, addr)

					if hook := cfg.AfterDialHook(); hook != nil {
						hook()
					}

					return conn, err
				},
				ForceAttemptHTTP2:     false,
				MaxIdleConns:          100,
				IdleConnTimeout:       3 * time.Second,
				TLSHandshakeTimeout:   1 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				DisableKeepAlives:     cfg.DisableKeepAlives(),
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if len(via) >= 5 {
					return http.ErrUseLastResponse
				}
				return nil
			},
			Jar:     nil,
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Get(_ context.Context, url string) ([]byte, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "c.client.Get")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "io.ReadAll")
	}

	return body, nil
}

func (c *Client) Do(_ context.Context, url string) (*http.Response, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "c.client.Get")
	}

	return resp, nil
}
