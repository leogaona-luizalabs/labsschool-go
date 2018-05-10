package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luizalabs/labsschool-go/api/handlers"
	"github.com/luizalabs/labsschool-go/config"
)

// StartServer ...
func StartServer() {

	// aqui criamos o router e em seguida registramos as rotas
	router := mux.NewRouter()
	// HandleFunc recebe um path e uma HandlerFunc, que por padrão tem a assinatura:
	// (w http.ResponseWriter, r *http.Request)
	// Por fim, setamos também os métodos que são aceitos nessa rota
	router.HandleFunc("/health", handlers.HealthCheck).Methods(http.MethodGet)
	router.HandleFunc("/api/users", handlers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users", handlers.ListUsers).Methods(http.MethodGet)
	http.Handle("/", router)

	// O server precisa de uma string que define em qual porta o servidor está escutar
	// A string fica no padrão ":PORTA"
	address := fmt.Sprintf(":%d", config.Port)

	// Aqui começamos a escutar as requisições na porta especificada
	http.ListenAndServe(address, nil)

}
