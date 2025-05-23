package models

type Transaction struct {
	ID           int     `json:"id"`
	Type         string  `json:"type"` // income, expense, transfer, loan
	Amount       float64 `json:"amount"`
	AccountFrom  int     `json:"account_from"`
	AccountTo    int     `json:"account_to"`
	CategoryID   int     `json:"category_id"`
	Description  string  `json:"description"`
	IsTaxApplied bool    `json:"is_tax_applied"`
	TaxAmount    float64 `json:"tax_amount"`
	Note         string  `json:"note"`
	Date         string  `json:"date"`
}
