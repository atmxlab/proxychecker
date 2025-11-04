package http

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/atmxlab/proxychecker/pkg/errors"
)

type Client struct {
	client *http.Client
}

func NewClient(proxyURL string) *Client {
	// TODO: cfg
	return &Client{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: func(r *http.Request) (*url.URL, error) {
					if proxyURL == "" { // TODO: remove
						return nil, nil
					}

					return url.Parse(proxyURL)
				},
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
	}
}

func (c *Client) Get(ctx context.Context, url string) ([]byte, error) {
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
