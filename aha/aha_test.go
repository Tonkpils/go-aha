package aha

import "testing"

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
