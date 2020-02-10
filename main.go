package main

import (
	"log"
	"net/http"
	"tensorflow-back/handlers"

	"github.com/rs/cors"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/user/login", handlers.LoginUser).Methods("POST")
	router.HandleFunc("/user", handlers.GetUser).Methods("GET")
	router.HandleFunc("/user", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/tensorflow", handlers.Tensorflow).Methods("GET")

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodPatch,
		},
		AllowedHeaders: []string{
			"Access-Control-Allow-Origin", "*",
		},
	})

	log.Printf("listening on port %s", "8081")
	http.ListenAndServe(":8081", corsOpts.Handler(router))
}
