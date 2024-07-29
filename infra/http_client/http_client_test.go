package http_client

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Get(t *testing.T) {
	// Cria um servidor de teste
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	client := &Client{HTTPClient: *server.Client()}

	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("esperava sem erro, mas obteve: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperava status %d, mas obteve %d", http.StatusOK, resp.StatusCode)
	}
}

func TestClient_Post(t *testing.T) {
	// Cria um servidor de teste
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Created"))
	}))
	defer server.Close()

	client := &Client{HTTPClient: *server.Client()}

	body := bytes.NewBufferString(`{"name":"test"}`)
	resp, err := client.Post(server.URL, "application/json", body)
	if err != nil {
		t.Fatalf("esperava sem erro, mas obteve: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("esperava status %d, mas obteve %d", http.StatusCreated, resp.StatusCode)
	}
}
