package models

// User represents a user with their relevant details.
type User struct {
	ID        int    `json:"id"`         // Unique identifier for the user
	Username  string `json:"username"`   // Email address of the user
	Password  string `json:"password"`   // Hashed password (not included in JSON output)
	CreatedAt string `json:"created_at"` // Timestamp when the user was created
}
