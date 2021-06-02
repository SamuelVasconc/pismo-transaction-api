package models

import "time"

//Account ...
type Account struct {
	ID                   int64   `json:"account_id,omitempty"`
	DocumentNumber       string  `json:"document_number,omitempty"`
	AvaliableCreditLimit float64 `json:"avaliable_credit_limit,omitempty"`
}

//Operation ...
type Operation struct {
	ID          int64  `json:"operation_type_id,omitempty"`
	Description string `json:"description,omitempty"`
}

//Transaction ...
type Transaction struct {
	ID              int64     `json:"transaction_id,omitempty"`
	AcountID        int64     `json:"account_id,omitempty"`
	OperationTypeID int64     `json:"operation_type_id,omitempty"`
	Amount          float64   `json:"amount,omitempty"`
	Balance         float64   `json:"balance,omitempty"`
	EventDate       time.Time `json:"event_date,omitempty"`
}
