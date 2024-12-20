package entities

import "time"

// Expense represents a single expense record
type Expense struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
}
