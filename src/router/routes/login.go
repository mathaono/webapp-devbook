package routes

import (
	"net/http"
	"webapp-devbook/src/controllers"
)

var loginRoutes = []Rota{
	{
		URI:               "/",
		Method:            http.MethodGet,
		Func:              controllers.LoadLoginScreen,
		ReqAuthentication: false,
	},
	{
		URI:               "/login",
		Method:            http.MethodGet,
		Func:              controllers.LoadLoginScreen,
		ReqAuthentication: false,
	},
}
