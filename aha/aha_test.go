package aha

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client, _ = NewClient(nil, "Test Account")
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	_, err := NewClient(nil, "")
	if err == nil {
		t.Errorf("NewClient expected err with no account name")
	}

	c, err := NewClient(nil, "Test Account")
	if err != nil {
		t.Errorf("NewClient expected err to be nil, got %v", err)
	}

	if c.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL is %v, expected %v", c.BaseURL.String(), defaultBaseURL)
	}

	if c.AccountName != "Test Account" {
		t.Errorf("NewClient AccountName is %v, expected %v", c.AccountName, "Test Account")
	}
}
