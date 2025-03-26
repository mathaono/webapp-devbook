package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Estrutura que representa todas as rotas da aplicação web
type Rota struct {
	URI               string
	Method            string
	Func              func(w http.ResponseWriter, r *http.Request)
	ReqAuthentication bool
}

// Inclui todas as rotas criadas dentro do router
func Configurate(router *mux.Router) *mux.Router {
	routes := loginRoutes

	for _, rota := range routes {
		router.HandleFunc(rota.URI, rota.Func).Methods(rota.Method)
	}

	return router
}
