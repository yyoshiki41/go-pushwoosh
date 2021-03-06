// Package pushwoosh provides functions and structs for accessing the Pushwoosh Remote API.
package pushwoosh

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	apiV13             = "1.3"
	defaultHTTPTimeout = 120 * time.Second
)

var (
	httpClient = &http.Client{Timeout: defaultHTTPTimeout}
)

// Config is a setting for Pushwoosh Remote APIs.
type Config struct {
	Endpoint        string
	ApplicationCode string
	AccessToken     string
}

// Result represents API Response for Pushwoosh.
type Result struct {
	StatusCode    int64            `json:"status_code"`
	StatusMessage string           `json:"status_message"`
	Response      ResponseMessages `json:"response"`
}

// ResponseMessages represents messages from Pushwoosh API
type ResponseMessages struct {
	Messages       []string            `json:"Messages"`
	UnknownDevices map[string][]string `json:"UnknownDevices"`
}

// Client represents an API client for Pushwoosh.
type Client struct {
	httpClient *http.Client
	config     *Config
}

// NewClient returns a new pushwoosh API client.
func NewClient(config *Config) (*Client, error) {
	if httpClient == nil {
		return nil, errors.New("httpClient is nil")
	}
	if config == nil {
		return nil, errors.New("config is nil")
	}

	return &Client{
		httpClient: httpClient,
		config:     config,
	}, nil
}

func (c *Client) call(ctx context.Context, method string, apiEndpoint string, params interface{}, res interface{}) error {
	u, err := url.Parse(c.config.Endpoint)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, apiV13, apiEndpoint)

	p, err := newRequestParams(c.config.ApplicationCode, c.config.AccessToken, params)
	if err != nil {
		return err
	}

	jsonParams, err := json.Marshal(p)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(jsonParams))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if res == nil {
		return nil
	}
	return json.NewDecoder(response.Body).Decode(&res)
}

type requestParams map[string]interface{}

func newRequestParams(application, auth string, params interface{}) (requestParams, error) {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var reqValues map[string]interface{}
	if err := json.Unmarshal(jsonParams, &reqValues); err != nil {
		return nil, err
	}
	reqValues["application"] = application
	reqValues["auth"] = auth

	return requestParams{
		"request": reqValues,
	}, nil
}

func (p *requestParams) setApplication(application string) {
	(*p)["application"] = application
}

func (p *requestParams) setAuth(auth string) {
	(*p)["auth"] = auth
}

// SetHTTPClient overrides the default HTTP client.
func SetHTTPClient(client *http.Client) {
	httpClient = client
}
