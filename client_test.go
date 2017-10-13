package pushwoosh

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := Config{
		Endpoint:        "https://cp.pushwoosh.com/json",
		ApplicationCode: "test-application-code",
		AccessToken:     "test-access-token",
	}

	_, err := NewClient(&c)
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}
}

func TestNewClient_EmptyHTTPClient(t *testing.T) {
	var emptyClient *http.Client

	SetHTTPClient(emptyClient)
	defer teardownHTTPClient()

	c := Config{
		Endpoint:        "https://cp.pushwoosh.com/json",
		ApplicationCode: "test-application-code",
		AccessToken:     "test-access-token",
	}
	client, err := NewClient(&c)
	if err == nil {
		t.Errorf("Should detect that HTTPClient is nil.\nclient: %v", client)
	}
}
