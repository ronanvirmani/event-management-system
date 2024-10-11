package models

// User represents a user structure
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Email    string `json:"email"`
}
