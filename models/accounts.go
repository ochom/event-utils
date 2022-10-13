package models

import (
	"log"
	"time"

	"github.com/dongri/phonenumber"
	"github.com/ochom/event-utils/auth"
	"gorm.io/gorm"
)

// Organization ...
type Organization struct {
	ID        string         `json:"id,omitempty"`
	ManagerID string         `json:"managerID,omitempty"`
	Name      string         `json:"name,omitempty"`
	About     string         `json:"about,omitempty"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
	Manager   User           `json:"manager,omitempty" gorm:"-"`
}

// User ...
type User struct {
	ID             string         `json:"id,omitempty"`
	OrganizationID string         `json:"organizationID,omitempty"`
	FirstName      string         `json:"firstName,omitempty"`
	LastName       string         `json:"lastName,omitempty"`
	Email          string         `json:"email,omitempty"`
	Image          string         `json:"image,omitempty"`
	Mobile         string         `json:"mobile,omitempty"`
	Password       string         `json:"-"`
	Group          UserGroup      `json:"group,omitempty"`
	CreatedAt      time.Time      `json:"createdAt,omitempty"`
	UpdatedAt      time.Time      `json:"updatedAt,omitempty"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt,omitempty"`
}

// Consumer the customers
type Consumer struct {
	ID          string         `json:"id,omitempty"`
	FirstName   string         `json:"firstName,omitempty"`
	LastName    string         `json:"lastName,omitempty"`
	Email       string         `json:"email,omitempty"`
	Image       string         `json:"image,omitempty"`
	Mobile      string         `json:"mobile,omitempty"`
	PIN         string         `json:"pin,omitempty"`
	IsNewPin    bool           `json:"isNewPin,omitempty"`
	DateOfBirth string         `json:"dob,omitempty"`
	CreatedAt   time.Time      `json:"createdAt,omitempty"`
	UpdatedAt   time.Time      `json:"updatedAt,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt,omitempty"`
}

// SetMobile sets the mobile number to kenyan mobile format 254 ...
func (c *Consumer) SetMobile(phone string) {
	phone = phonenumber.Parse(phone, "KE")
	c.Mobile = phone
}

// SetPIN hash pin and update
func (c *Consumer) SetPIN(pin string) {
	hashed, err := auth.HashPassword(pin)
	if err != nil {
		log.Println("auth.HashPassword: ", err.Error())
		return
	}

	c.PIN = *hashed
}

// DashboardSummary ...
type DashboardSummary struct {
	TotalEvents  int `json:"totalEvents,omitempty"`
	TicketsSold  int `json:"ticketsSold,omitempty"`
	TicketsLeft  int `json:"ticketsLeft,omitempty"`
	TotalRevenue int `json:"totalRevenue,omitempty"`
}

// DashboardEvent ...
type DashboardEvent struct {
	EventID     string `json:"eventID,omitempty"`
	Name        string `json:"name,omitempty"`
	TicketsSold int    `json:"ticketsSold,omitempty"`
	Revenue     int    `json:"revenue,omitempty"`
}

// EventSummary ...
type EventSummary struct {
	EventID      string `json:"eventID,omitempty"`
	Name         string `json:"name,omitempty"`
	TotalTickets int    `json:"totalTickets,omitempty"`
	TicketsSold  int    `json:"ticketsSold,omitempty"`
	TicketsLeft  int    `json:"ticketsLeft,omitempty"`
	Revenue      int    `json:"revenue,omitempty"`
}

// EventReport ...
type EventReport struct {
	EventID      string    `json:"eventID,omitempty"`
	PaymentID    string    `json:"paymentID,omitempty"`
	ConsumerName string    `json:"consumerName,omitempty"`
	TicketName   string    `json:"ticketName,omitempty"`
	Quantity     int       `json:"quantity,omitempty"`
	Price        int       `json:"price,omitempty"`
	Status       string    `json:"status,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
}

// TicketSummary ...
type TicketSummary struct {
	Name    string `json:"ticketName,omitempty"`
	Total   int    `json:"totalTickets,omitempty"`
	Sold    int    `json:"ticketsSold,omitempty"`
	Left    int    `json:"ticketsLeft,omitempty"`
	Revenue int    `json:"revenue,omitempty"`
}

// EventTicketSummary ...
type EventTicketSummary struct {
	Event          Event           `json:"event,omitempty"`
	TicketsSummary []TicketSummary `json:"ticketsSummary,omitempty"`
}
