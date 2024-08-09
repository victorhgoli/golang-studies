package controller

import (
	logger "estudo-test/infra/logger"
	"estudo-test/internal/service"
	"estudo-test/pkg/kafka/producer"
	"log"
	"net/http"
)

type AsyncCadRequest struct {
	Nome    string `json:"nome"`
	Email   string `json:"email"`
	Produto string `json:"produto"`
}

type asyncCadController struct {
	UserService   service.UserService
	PedidoService service.PedidoService
	Log           logger.Logger
}

type AsyncCadController interface {
	Cadastrar(w http.ResponseWriter, r *http.Request)
}

func NewAsyncCadController(userService service.UserService, pedidoService service.PedidoService, log logger.Logger) AsyncCadController {
	return &asyncCadController{
		UserService:   userService,
		PedidoService: pedidoService,
		Log:           log,
	}
}

func (c *asyncCadController) Cadastrar(w http.ResponseWriter, r *http.Request) {
	p, err := producer.NewProducer()
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}

	err = p.ProduceMessage()
	if err != nil {
		log.Fatalf("Error producing message: %v", err)
	}
}
