package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// logger := util.GetLogger().WithFields(map[string]interface{}{
	// 	"module":        "main",
	// 	"function_name": "main",
	// })

	// dsn := "root:root@tcp(localhost:3306)/optimusprime"

	// db, err := util.OpenMySQLConnection(dsn)
	// if err != nil {
	// 	logger.WithField("error", err.Error()).Error()
	// 	os.Exit(1)
	// }

	// service := routes.New(db)
	// items, err := service.List(50, 0)
	// if err != nil {
	// 	logger.WithField("error", err.Error()).Error()
	// 	os.Exit(1)
	// }

	// result := map[string]interface{}{
	// 	"routes": items,
	// }

	// responseBytes, err := json.Marshal(&result)
	// if err != nil {
	// 	logger.WithField("error", err.Error()).Error()
	// 	os.Exit(1)
	// }

	// fmt.Println(string(responseBytes))

	// unmarshallRoutes := routes.Routes{}
	// err := json.Unmarshal([]byte(routesJSON), &unmarshallRoutes)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	// fmt.Println(unmarshallRoutes[0].ID)

	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("POST")

	r.Use(autenticador)

	http.ListenAndServe(":9876", r)
}

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusCreated)

	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	w.Write(resp)
}

func autenticador(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")
		if token == "xispirito" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
