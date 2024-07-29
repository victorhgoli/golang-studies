package service

import (
	"estudo-test/internal/repository"
	"estudo-test/pkg/models"
)

type pedidoService struct {
	Repo repository.PedidoRepository
}

type PedidoService interface {
	CreatePedido(produto string, user_id int64) (int64, error)
}

func NewPedidoService(repo repository.PedidoRepository) PedidoService {
	return &pedidoService{Repo: repo}
}

func (s *pedidoService) CreatePedido(produto string, user_id int64) (int64, error) {
	pedido := &models.Pedido{Produto: produto, UserId: user_id}
	id, er := s.Repo.InsertPedido(pedido)

	return id, er
}
