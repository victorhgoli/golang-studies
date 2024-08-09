package routes

import (
	"estudo-test/api/controller"
	"estudo-test/api/graphql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/graphql-go/handler"
)

func NewRouter(cadController controller.CadController) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/cadastrar", cadController.Cadastrar)
	r.Get("/usuario/{id}", cadController.Buscar)

	h := handler.New(&handler.Config{
		Schema:   &graphql.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	r.Handle("/graphql", h)

	return r
}
