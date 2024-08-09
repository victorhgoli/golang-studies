package integration

import (
	"encoding/json"
	"estudo-test/infra/logger"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/sony/gobreaker"
)

type infoTestIntegration struct {
	Log            logger.Logger
	HTTP           *http.Client
	CircuitBreaker *gobreaker.CircuitBreaker
}

func NewInfoTestIntegration(log logger.Logger) InfoTestIntegration {
	circuit := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "GetInfoTest",
		MaxRequests: 3,
		Interval:    60 * time.Second,
		Timeout:     30 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			log.Infof("counts %s", int(counts.ConsecutiveFailures))
			return counts.ConsecutiveFailures > 3
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			log.Infof("Circuit breaker state changed from %v to %v", from, to)
		},
	})

	return &infoTestIntegration{
		Log:            log,
		HTTP:           &http.Client{},
		CircuitBreaker: circuit,
	}
}

func (i *infoTestIntegration) GetInfo() (interface{}, error) {
	url := "https://jsonplaceholder.typicode.com7/todos/1"

	var data map[string]interface{}
	operation := func() error {
		response, err := i.HTTP.Get(url)
		if err != nil {
			i.Log.Errorf("Erro ao fazer a requisição: %v", err)
			return err
		}
		defer response.Body.Close()

		if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
			i.Log.Errorf("Erro ao decodificar a resposta: %v", err)
			return err
		}

		return nil
	}

	body, err := i.CircuitBreaker.Execute(func() (interface{}, error) {
		expBackoff := backoff.NewExponentialBackOff()
		expBackoff.MaxElapsedTime = 15 * time.Second
		retryBackoff := backoff.WithMaxRetries(expBackoff, 3)

		err := backoff.Retry(operation, retryBackoff)

		if err != nil {
			i.Log.Errorf("Erro após múltiplas tentativas: %v", err)
			return nil, err
		}

		return data, nil
	})

	if err != nil {
		return nil, err
	}

	return body, nil
}
