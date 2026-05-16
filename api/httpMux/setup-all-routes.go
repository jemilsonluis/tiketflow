package httpmux

import (
	"net/http"

	"github.com/jemilsonluis/api/handlers"
	"github.com/jemilsonluis/internals/modules/users/factory"
	"github.com/jemilsonluis/internals/modules/users/infra/repository"
)

func SetupAllRoutes(mux *http.ServeMux) {
	repo := repository.NewUserRepositoryImpl()
	factory := factory.NewUserFactoryImpl(repo)
	handler := handlers.NewUserhandlerImpl(factory)

	SetupUserRoutes(mux, handler)
}
