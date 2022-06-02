package models

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/datatypes"
)

// Event ...
type Event struct {
	ID             string         `json:"id,omitempty"`
	OrganizationID string         `json:"organizationID,omitempty"`
	Name           string         `json:"name,omitempty"`
	Image          string         `json:"image,omitempty"`
	Category       string         `json:"category,omitempty"`
	Description    string         `json:"description,omitempty"`
	Venue          string         `json:"venue,omitempty"`
	StartTime      time.Time      `json:"startTime,omitempty"`
	EndTime        time.Time      `json:"endTime,omitempty"`
	Facebook       string         `json:"facebook,omitempty"`
	Twitter        string         `json:"twitter,omitempty"`
	HashTags       string         `json:"hashTags,omitempty"`
	TnC            string         `json:"tnc,omitempty"`
	Tickets        datatypes.JSON `json:"tickets,omitempty"`
	CreatedAt      time.Time      `json:"createdAt,omitempty"`
	UpdatedAt      time.Time      `json:"updatedAt,omitempty"`
}

// SetTickets ...
func (e *Event) SetTickets(data []EventTicket) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	e.Tickets = b

	return nil
}

// GetTickets ...
func (e *Event) GetTickets() ([]*EventTicket, error) {
	data := []*EventTicket{}
	if err := json.Unmarshal(e.Tickets, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// GetTicket gets one ticket by name
func (e *Event) GetTicket(name string) (*EventTicket, error) {
	tickets, err := e.GetTickets()
	if err != nil {
		return nil, err
	}

	for _, v := range tickets {
		if v.Name == name {
			return v, nil
		}
	}

	return nil, fmt.Errorf("ticket `%s` not found", name)
}

// EventTicket ...
type EventTicket struct {
	Name     string    `json:"name,omitempty"`
	Cost     int       `json:"cost,omitempty"`
	Slots    int       `json:"slots,omitempty"`
	Deadline time.Time `json:"deadline,omitempty"`
}

// Payment ...
type Payment struct {
	ID         string    `json:"id,omitempty"`
	ConsumerID string    `json:"consumerID,omitempty"`
	Mobile     string    `json:"mobile,omitempty"`
	EventID    string    `json:"eventID,omitempty"`
	TicketName string    `json:"ticketID,omitempty"`
	Quantity   int       `json:"quantity,omitempty"`
	Amount     int       `json:"amount,omitempty"`
	Status     string    `json:"status,omitempty"`
	MerchantID string    `json:"merchantID,omitempty"`
	Paid       bool      `json:"paid,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

// Booking ...
type Booking struct {
	ID         string    `json:"id,omitempty"`
	EventID    string    `json:"eventID,omitempty"`
	TicketName string    `json:"ticketID,omitempty"`
	ConsumerID string    `json:"consumerID,omitempty"`
	PaymentID  string    `json:"paymentID,omitempty"`
	Number     string    `json:"number,omitempty"`
	Quantity   int       `json:"quantity,omitempty"`
	Amount     int       `json:"amount,omitempty"`
	Used       bool      `json:"used,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}
