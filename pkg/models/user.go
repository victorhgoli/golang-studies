package models

type User struct {
	ID    int64
	Nome  string
	Email string
}

type Pedido struct {
	ID      int64
	Produto string
	UserId  int64
}
