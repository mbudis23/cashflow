package models

type APIMessage struct {
	Message string `json:"message"`
}

type APIError struct {
	Error string `json:"error"`
}
