package httpmux

import (
	"net/http"

)

func StartServerHttpMux() *http.ServeMux {
	mux := http.NewServeMux()
	SetupAllRoutes(mux)

	return mux
}
