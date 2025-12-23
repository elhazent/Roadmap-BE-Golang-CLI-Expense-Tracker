package model

import "time"

type ExpenseModel struct {
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Category    string    `json:"category,omitempty"`
}

type ExpenseData struct {
	ExpenseData []ExpenseModel `json:"data"`
}