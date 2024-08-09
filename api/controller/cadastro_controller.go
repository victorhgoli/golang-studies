package controller

import (
	"context"
	"encoding/json"
	logger "estudo-test/infra/logger"
	"estudo-test/internal/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type CadRequest struct {
	Nome    string `json:"nome"`
	Email   string `json:"email"`
	Produto string `json:"produto"`
}

type cadController struct {
	UserService   service.UserService
	PedidoService service.PedidoService
	Log           logger.Logger
}

type CadController interface {
	Cadastrar(w http.ResponseWriter, r *http.Request)
	Buscar(w http.ResponseWriter, r *http.Request)
}

func NewCadController(userService service.UserService, pedidoService service.PedidoService, log logger.Logger) CadController {
	return &cadController{
		UserService:   userService,
		PedidoService: pedidoService,
		Log:           log,
	}
}

func (c *cadController) Cadastrar(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var userReq CadRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	done := make(chan struct{})
	var id, prodId int64
	var err error

	go func() {
		defer close(done)
		// Inserir um registro na tabela usuario
		id, prodId, err = c.createUserAndPedido(userReq)
	}()

	select {
	case <-ctx.Done():
		http.Error(w, "Tempo de requisição esgotado", http.StatusGatewayTimeout)
		fmt.Println("Contexto cancelado:", ctx.Err())
	case <-done:
		if err != nil {
			// Erro já tratado dentro da goroutine
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Usuário e Pedido criados com sucesso! UserID: %d, PedidoID: %d\n", id, prodId)
	}
}

func (c *cadController) createUserAndPedido(userReq CadRequest) (int64, int64, error) {
	id, err := c.UserService.CreateUser(userReq.Nome, userReq.Email)
	if err != nil {
		log.Println("Erro ao criar usuário:", err)
		return 0, 0, err
	}

	prodId, err := c.PedidoService.CreatePedido(userReq.Produto, id)
	if err != nil {
		log.Println("Erro ao criar pedido:", err)
		return 0, 0, err
	}

	return id, prodId, nil
}

func (c *cadController) Buscar(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	user, err := c.UserService.GetUsersById(id)
	if err != nil {
		http.Error(w, "Erro ao buscar usuário", http.StatusInternalServerError)
		c.Log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(user)

}
