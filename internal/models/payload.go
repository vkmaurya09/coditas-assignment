package models

type Payload struct {
	Name   string `json:"name" validate:"required"`
	PAN    string `json:"pan" validate:"required,pan"`
	Mobile string `json:"mobile" validate:"required,len=10,numeric"`
	Email  string `json:"email" validate:"required,email"`
}
