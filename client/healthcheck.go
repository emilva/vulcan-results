// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "vulcan-results": healthcheck Resource Client
//
// Command:
// $ goagen
// --design=github.com/adevinta/vulcan-results/design
// --out=$(GOPATH)/src/github.com/adevinta/vulcan-results
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ShowHealthcheckPath computes a request path to the show action of healthcheck.
func ShowHealthcheckPath() string {

	return fmt.Sprintf("/healthcheck")
}

// Get the health status for the application
func (c *Client) ShowHealthcheck(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowHealthcheckRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowHealthcheckRequest create the request corresponding to the show action endpoint of the healthcheck resource.
func (c *Client) NewShowHealthcheckRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}