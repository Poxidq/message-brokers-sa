package models

// Message struct to match the structure of the JSON message
type Message struct {
	Message string `json:"message"`
	User    string `json:"user"`
}
