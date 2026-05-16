package bootstrap

import (
	"fmt"
	"net/http"
)

func Bootstrap(server http.Handler) {
	fmt.Println("server running in port 8000")
	http.ListenAndServe(":8000", server)
}
