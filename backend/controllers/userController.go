package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/ronanvirmani/event-management-system/backend/models"
    "github.com/ronanvirmani/event-management-system/backend/services"
)

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    err := services.SignUpUser(user)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(err.Error())
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode("User registered successfully")
}

// LoginUser handles user login
func LoginUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    token, err := services.LoginUser(user)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(err.Error())
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
