package repository

import (
	"database/sql"
	"estudo-test/pkg/models"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) InsertUser(user *models.User) (int64, error) {

	query := "INSERT INTO usuario (nome, email) VALUES (?, ?)"

	result, err := r.DB.Exec(query, user.Nome, user.Email)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *userRepository) GetUserById(id int64) (*models.User, error) {

	query := `SELECT id, nome, email FROM usuario WHERE ID = ?`

	row := r.DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Nome, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found with the given ID
		}
		return nil, err
	}

	return &user, nil
}
