package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ronanvirmani/event-management-system/backend/routes"
)

func main() {

    r := mux.NewRouter()
    routes.RegisterRoutes(r)

    // Allow CORS
    headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
    origins := handlers.AllowedOrigins([]string{"*"})

    log.Println("Server listening on port 8000")
    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(r)))
}
