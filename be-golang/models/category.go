package models

type Category struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id,omitempty"`
	Name   string `json:"name"`
	Color  string `json:"color"` // Format: HEX atau nama warna (misal: "#FFAA00" atau "blue")
}
