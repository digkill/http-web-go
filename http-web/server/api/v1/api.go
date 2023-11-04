package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strings"
)

const ValidBearer = "123456"

type MessageResponse struct {
	Message string `json:"message"`
}

func jsonResponse(w http.ResponseWriter, data interface{}, c int) {
	output, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", output)
}

func Hey(w http.ResponseWriter, r *http.Request) {
	response := MessageResponse{
		Message: "Hey!",
	}
	jsonResponse(w, response, http.StatusOK)
}

func HeyByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	response := MessageResponse{
		Message: fmt.Sprintf("Hey %s!", name),
	}
	jsonResponse(w, response, http.StatusOK)
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Authorization")
		if token == "" {
			http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		if token != ValidBearer {
			http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(writer, request)
	})

}

func NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(AuthenticationMiddleware)

	router.Get("/", Hey)
	router.Get("/{name}", HeyByName)

	return router
}
