package models

import (
	"time"
)

// Payment ...
type Payment struct {
	ID         string    `json:"id,omitempty"`
	ConsumerID string    `json:"consumerID,omitempty"`
	EventID    string    `json:"eventID,omitempty"`
	TicketName string    `json:"ticketID,omitempty"`
	Quantity   int       `json:"quantity,omitempty"`
	Amount     int       `json:"amount,omitempty"`
	Paid       bool      `json:"paid,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

// Booking ...
type Booking struct {
	ID         string    `json:"id,omitempty"`
	EventID    string    `json:"eventID,omitempty"`
	TicketID   string    `json:"ticketID,omitempty"`
	ConsumerID string    `json:"consumerID,omitempty"`
	PaymentID  string    `json:"paymentID,omitempty"`
	Code       string    `json:"code,omitempty"`
	QRCode     string    `json:"qrCode,omitempty"`
	Quantity   int       `json:"quantity,omitempty"`
	Amount     int       `json:"amount,omitempty"`
	Used       bool      `json:"used,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}
