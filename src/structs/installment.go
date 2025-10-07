package structs

import "github.com/google/uuid"

type Installment struct {
	ID           uuid.UUID `json:"id"`
	ValueInCents int       `json:"value_in_cents"`
	DueDay       int       `json:"due_day"`
}
