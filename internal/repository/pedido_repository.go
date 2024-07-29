package repository

import (
	"database/sql"
	"estudo-test/pkg/models"
)

type pedidoRepository struct {
	DB *sql.DB
}

func NewPedidoRepository(db *sql.DB) PedidoRepository {
	return &pedidoRepository{DB: db}
}

func (r *pedidoRepository) InsertPedido(user *models.Pedido) (int64, error) {

	query := "INSERT INTO pedido (produto, user_id) VALUES (?, ?)"

	result, err := r.DB.Exec(query, user.Produto, user.UserId)

	if err != nil {
		return 0, nil
	}

	return result.LastInsertId()
}
