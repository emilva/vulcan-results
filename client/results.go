// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "vulcan-results": Results Resource Client
//
// Command:
// $ goagen
// --design=github.com/adevinta/vulcan-results/design
// --out=$(GOPATH)/src/github.com/adevinta/vulcan-results
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// RawResultsPath computes a request path to the raw action of Results.
func RawResultsPath() string {

	return fmt.Sprintf("/v1/raw")
}

// Update the Raw of a Check
func (c *Client) RawResults(ctx context.Context, path string, payload *RawPayload) (*http.Response, error) {
	req, err := c.NewRawResultsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRawResultsRequest create the request corresponding to the raw action endpoint of the Results resource.
func (c *Client) NewRawResultsRequest(ctx context.Context, path string, payload *RawPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// ReportResultsPath computes a request path to the report action of Results.
func ReportResultsPath() string {

	return fmt.Sprintf("/v1/report")
}

// Update the Report of a Check
func (c *Client) ReportResults(ctx context.Context, path string, payload *ReportPayload) (*http.Response, error) {
	req, err := c.NewReportResultsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewReportResultsRequest create the request corresponding to the report action endpoint of the Results resource.
func (c *Client) NewReportResultsRequest(ctx context.Context, path string, payload *ReportPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}
