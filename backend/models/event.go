package models

type Event struct {
	ID 		string `json:"id"`
	Title 	string `json:"title"`
	Description string `json:"description"`
	Location string `json:"location"`
	Date string `json:"date"`
	ImageURL string `json:"imageURL,omitempty"`
}