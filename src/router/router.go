package router

import (
	"webapp-devbook/src/router/routes"

	"github.com/gorilla/mux"
)

// Retorna um Router com todas as rotas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configurate(r)
}
