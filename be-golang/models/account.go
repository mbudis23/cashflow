package models

type Account struct {
	ID      int     `json:"id"`
	UserID  int     `json:"user_id,omitempty"`
	Name    string  `json:"name"`
	Type    string  `json:"type"`    // 'cash', 'bank', 'ewallet'
	Balance float64 `json:"balance"` // Default = 0
}
