package api

import (
	"net/http"

	"portfolio"
)

var handler = portfolio.NewHandler()

// Handler is the Vercel Go runtime entrypoint; vercel.json rewrites all
// paths here.
func Handler(w http.ResponseWriter, r *http.Request) {
	handler.ServeHTTP(w, r)
}
