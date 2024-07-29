package repository

import "estudo-test/pkg/models"

type PedidoRepository interface {
	InsertPedido(user *models.Pedido) (int64, error)
}

type UserRepository interface {
	InsertUser(user *models.User) (int64, error)
	GetUserById(id int64) (*models.User, error)
}
