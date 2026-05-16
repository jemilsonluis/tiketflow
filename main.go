package main

import (
	"github.com/jemilsonluis/api/bootstrap"
	httpmux "github.com/jemilsonluis/api/httpMux"
)

func main() {
	server := httpmux.StartServerHttpMux()
	bootstrap.Bootstrap(server)
}
