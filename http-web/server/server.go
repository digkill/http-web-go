package server

import (
	"fmt"
	"github.com/go-chi/chi"
	v1 "http-web/server/api/v1"
	"net/http"
	"path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func NewRouter() http.Handler {
	router := chi.NewRouter()
	router.Get("/", Home)

	router.Mount("/api/v1/", v1.NewRouter())

	staticPath, _ := filepath.Abs("../../static")
	fs := http.FileServer(http.Dir(staticPath))
	router.Handle("/*", fs)

	return router
}
