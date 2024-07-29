package http_client

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

// HTTPClient é uma interface para facilitar o mock em testes
type HTTPClient interface {
	Get(url string) (*http.Response, error)
	Post(url, contentType string, body *bytes.Buffer) (*http.Response, error)
}

// Client é uma implementação concreta de HTTPClient
type Client struct {
	HTTPClient http.Client
}

// NewClient cria uma nova instância de Client
func NewClient() HTTPClient {
	return &Client{HTTPClient: http.Client{}}
}

// Get faz uma requisição HTTP GET e registra a duração da chamada
func (c *Client) Get(url string) (*http.Response, error) {
	start := time.Now()
	response, err := c.HTTPClient.Get(url)
	duration := time.Since(start)

	if err != nil {
		log.Printf("Erro ao fazer a requisição GET: %v", err)
		return nil, err
	}

	log.Printf("Requisição GET para %s levou %v", url, duration)
	return response, nil
}

// Post faz uma requisição HTTP POST e registra a duração da chamada
func (c *Client) Post(url, contentType string, body *bytes.Buffer) (*http.Response, error) {
	start := time.Now()
	response, err := c.HTTPClient.Post(url, contentType, body)
	duration := time.Since(start)

	if err != nil {
		log.Printf("Erro ao fazer a requisição POST: %v", err)
		return nil, err
	}

	log.Printf("Requisição POST para %s levou %v", url, duration)
	return response, nil
}
