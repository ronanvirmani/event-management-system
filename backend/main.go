package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ronanvirmani/event-management-system/backend/routes"
)

func main() {
	godotenv.Load(".env")

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	http.Handle("/", r)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}