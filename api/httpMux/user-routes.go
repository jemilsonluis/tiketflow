package httpmux

import (
	"fmt"
	"net/http"

	"github.com/jemilsonluis/api/configs"
	"github.com/jemilsonluis/api/handlers"
)

func SetupUserRoutes(server *http.ServeMux, handler handlers.IUserHandler) {
	fmt.Println("POST " + configs.UserRoutes.Create)
	server.HandleFunc("POST "+configs.UserRoutes.Create, handler.Create)
	server.HandleFunc("DELETE "+configs.UserRoutes.Delete, handler.Delete)
	server.HandleFunc("GET "+configs.UserRoutes.Fetch, handler.Fetch)
	server.HandleFunc("PATCH "+configs.UserRoutes.FindByEmail, handler.FetchByEmail)
	server.HandleFunc("GET "+configs.UserRoutes.Find, handler.Find)
}
