package routes

import (
	"github.com/gorilla/mux"
	"github.com/ronanvirmani/event-management-system/backend/controllers"
)

func RegisterRoutes(r *mux.Router) {

	// event routes
	r.HandleFunc("/events", controllers.GetEvents).Methods("GET")
	r.HandleFunc("/events", controllers.CreateEvent).Methods("POST")
	r.HandleFunc("/events/{id}", controllers.GetEvent).Methods("GET")
	r.HandleFunc("/events/{id}", controllers.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", controllers.DeleteEvent).Methods("DELETE")

	// user routes
	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")
}