package database

import (
	"fmt"

	"github.com/ochom/event-utils/models"
	"github.com/ochom/event-utils/utils"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// impl implements the database using postgres
type impl struct {
	db *gorm.DB
}

// DBKind database kind e.g postgres, sqlite
type DBKind string

const (
	// Postgres is the postgres database
	Postgres DBKind = "postgres"
	// SQLite is the sqlite database
	SQLite DBKind = "sqlite"
)

const triggerQuery = `
CREATE OR REPLACE FUNCTION count_tickets() 
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
AS 
$$
  DECLARE
    total int;
  BEGIN
    SELECT count(*) + 1  INTO total FROM bookings t WHERE t.event_id=NEW.event_id AND t.ticket_name=NEW.ticket_name;
    NEW.ticket_number := total;
    return NEW;
  END;
$$;

DROP TRIGGER IF EXISTS count_tickets_trigger
  ON public.bookings;

CREATE TRIGGER count_tickets_trigger 
  BEFORE INSERT on public.bookings
  FOR EACH ROW
  EXECUTE PROCEDURE count_tickets();
`

// Repository ...
type Repository interface {
	CreateOrUpdateOrganization(data *models.Organization) error
	DeleteOrganization(query *models.Organization) error
	GetOrganization(query *models.Organization) (*models.Organization, error)
	GetOrganizations(query *models.Organization) ([]*models.Organization, error)

	CreateOrUpdateUser(data *models.User) error
	DeleteUser(query *models.User) error
	GetUser(query *models.User) (*models.User, error)
	GetUsers(query *models.User) ([]*models.User, error)

	CreateOrUpdateConsumer(data *models.Consumer) error
	DeleteConsumer(query *models.Consumer) error
	GetConsumer(query *models.Consumer) (*models.Consumer, error)
	GetConsumers(query *models.Consumer) ([]*models.Consumer, error)

	CreateOrUpdateEvent(data *models.Event) error
	DeleteEvent(query *models.Event) error
	GetEvent(query *models.Event) (*models.Event, error)
	GetEvents(query *models.Event) ([]*models.Event, error)
	GetActiveEvents() ([]*models.Event, error)

	CreateOrUpdatePayment(data *models.Payment) error
	DeletePayment(query *models.Payment) error
	GetPayment(query *models.Payment) (*models.Payment, error)
	GetPayments(query *models.Payment) ([]*models.Payment, error)
	GetDistinctPayments() ([]*EventGroup, error)

	CreateOrUpdateBooking(data *models.Booking) error
	DeleteBooking(query *models.Booking) error
	GetBooking(query *models.Booking) (*models.Booking, error)
	GetBookings(query *models.Booking) ([]*models.Booking, error)
}

// New creates a new Database instance for repository
func New(kind DBKind) (Repository, error) {

	db, err := &gorm.DB{}, error(nil)

	switch kind {
	case Postgres:
		dns := utils.MustGetEnv("DATABASE_DNS")
		db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		break
	default:
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		break
	}

	err = db.AutoMigrate(
		&models.Organization{},
		&models.Event{},
		&models.User{},
		&models.Consumer{},
		&models.Payment{},
		&models.Booking{},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	// create trigger
	if kind == Postgres {
		if err := db.Exec(triggerQuery).Error; err != nil {
			return nil, fmt.Errorf("failed to create trigger: %w", err)
		}
	}

	return &impl{db}, nil
}
