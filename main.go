package main

import (
	"log"
	"net/http"
	"tes-rssa/database"
	"tes-rssa/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/user/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user", handlers.GetAllUser).Methods("GET")
	r.HandleFunc("/create", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/update", handlers.UpdateUser).Methods("POST")

	log.Println("Server berjalan di port 8002")
	log.Fatal(http.ListenAndServe(":8002", r))
}
