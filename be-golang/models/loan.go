package models

type Loan struct {
	ID                   int     `json:"id"`
	UserID               int     `json:"user_id,omitempty"`
	ContactName          string  `json:"contact_name"`
	Amount               float64 `json:"amount"`
	LoanType             string  `json:"loan_type"` // "given" atau "received"
	DueDate              string  `json:"due_date"`
	Interest             float64 `json:"interest"`
	IsPaid               bool    `json:"is_paid"`
	Note                 string  `json:"note"`
	RelatedTransactionID int     `json:"related_transaction_id"`
}
