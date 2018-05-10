package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luizalabs/labsschool-go/db"
	"github.com/luizalabs/labsschool-go/model"
)

// ListUsers handler da rota GET /users. Funcionamento parecido com o da rota GET /health
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	response, _ := json.Marshal(map[string]interface{}{
		"users": db.Users,
	})

	w.Write(response)
	w.WriteHeader(http.StatusOK)
	return
}

// CreateUser handler da rota POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Body é do tipo ReadCloser, ou seja, precisa ser fechado depois que for utilizado
	// Para isso, o ideal é usarmos um defer
	defer r.Body.Close()
	w.Header().Set("content-type", "application/json")

	// Instanciamos um User para ter um endereço de memória alocado para a struct
	requestUser := model.User{}
	// Parse do body na struct. Sempre passem um ponteiro e não uma cópia da struct!
	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(map[string]interface{}{
			"code":       500,
			"message":    err.Error(),
			"error_type": "parse request body",
		})
		w.Write(response)
		return
	}

	// Preenche o id e faz o append do user atual no "banco de dados"
	requestUser.ID = len(db.Users) + 1
	db.Users = append(db.Users, requestUser)

	// Devolvemos o User com ID preenchido no corpo da resposta
	response, _ := json.Marshal(requestUser)
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	return
}
