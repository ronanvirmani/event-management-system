package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/ronanvirmani/event-management-system/backend/models"
    "github.com/ronanvirmani/event-management-system/backend/services"
)

var events []models.Event

func init() {
    events = []models.Event{}
}

// GetEvents retrieves all events
func GetEvents(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(events)
}

// CreateEvent adds a new event
func CreateEvent(w http.ResponseWriter, r *http.Request) {
    var event models.Event
    _ = json.NewDecoder(r.Body).Decode(&event)
    event.ID = services.GenerateID()
    events = append(events, event)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(event)
}

// GetEvent retrieves a single event by ID
func GetEvent(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range events {
        if item.ID == params["id"] {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode("Event not found")
}

// UpdateEvent updates an existing event
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range events {
        if item.ID == params["id"] {
            events = append(events[:index], events[index+1:]...)
            var event models.Event
            _ = json.NewDecoder(r.Body).Decode(&event)
            event.ID = params["id"]
            events = append(events, event)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(event)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode("Event not found")
}

// DeleteEvent removes an event
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range events {
        if item.ID == params["id"] {
            events = append(events[:index], events[index+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode("Event not found")
}
