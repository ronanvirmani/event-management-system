package routes

import (
    "github.com/gorilla/mux"
    "github.com/ronanvirmani/event-management-system/backend/controllers"
    "github.com/ronanvirmani/event-management-system/backend/middleware"
)

func RegisterRoutes(r *mux.Router) {
    // Public routes
    r.HandleFunc("/api/events", controllers.GetEvents).Methods("GET")
    r.HandleFunc("/api/events/{id}", controllers.GetEvent).Methods("GET")
    r.HandleFunc("/api/register", controllers.RegisterUser).Methods("POST")
    r.HandleFunc("/api/login", controllers.LoginUser).Methods("POST")

    // Protected routes
    s := r.PathPrefix("/api").Subrouter()
    s.Use(middleware.AuthMiddleware)
    s.HandleFunc("/events", controllers.CreateEvent).Methods("POST")
    s.HandleFunc("/events/{id}", controllers.UpdateEvent).Methods("PUT")
    s.HandleFunc("/events/{id}", controllers.DeleteEvent).Methods("DELETE")
    s.HandleFunc("/upload", controllers.UploadFile).Methods("POST")
}
