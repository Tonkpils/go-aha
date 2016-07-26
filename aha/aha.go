package aha

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://secure.aha.io/api/v1/"
	userAgent      = "go-aha/1"
)

type Client struct {
	client *http.Client

	BaseURL   *url.URL
	UserAgent string

	AccountName string
	Username    string
	Password    string

	Products *ProductsService
}

func NewClient(client *http.Client, accountName string) (*Client, error) {
	if client == nil {
		client = http.DefaultClient
	}

	if accountName == "" {
		return nil, errors.New("Aha! account name is required")
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:      client,
		BaseURL:     baseURL,
		AccountName: accountName,
	}
	c.Products = &ProductsService{client: c}

	return c, nil
}

func NewBasicAuthClient(username, password, accountName string) (*Client, error) {
	client, err := NewClient(nil, accountName)
	if err != nil {
		return nil, err
	}

	client.Username = username
	client.Password = password

	return client, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Aha-Account", c.AccountName)
	req.Header.Set("Accept", "application/json")
	if c.Username != "" && c.Password != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
