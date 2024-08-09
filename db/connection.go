package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verificar se a conexão está funcionando
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Conectado ao banco de dados!")
	return db, nil
}
