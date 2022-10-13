package database

import (
	"context"
	"time"

	"github.com/ochom/event-utils/models"
	"github.com/ochom/event-utils/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// impl implements the database using postgres
type impl struct {
	db *gorm.DB
}

func initDB() (*gorm.DB, error) {
	dns := utils.MustGetEnv("DATABASE_DNS")

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	return db, err
}

func migrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Organization{},
		&models.Event{},
		&models.User{},
		&models.Consumer{},
		&models.Payment{},
		&models.Booking{},
	)

	return err
}

// Repository ...
type Repository interface {
	CreateOrUpdateOrganization(ctx context.Context, data models.Organization) error
	DeleteOrganization(ctx context.Context, query models.Organization) error
	GetOrganization(ctx context.Context, query models.Organization) (*models.Organization, error)
	GetOrganizations(ctx context.Context, query models.Organization) ([]*models.Organization, error)

	CreateOrUpdateUser(ctx context.Context, data models.User) error
	DeleteUser(ctx context.Context, query models.User) error
	GetUser(ctx context.Context, query models.User) (*models.User, error)
	GetUsers(ctx context.Context, query models.User) ([]*models.User, error)

	CreateOrUpdateConsumer(ctx context.Context, data models.Consumer) error
	DeleteConsumer(ctx context.Context, query models.Consumer) error
	GetConsumer(ctx context.Context, query models.Consumer) (*models.Consumer, error)
	GetConsumers(ctx context.Context, query models.Consumer) ([]*models.Consumer, error)

	CreateOrUpdateEvent(ctx context.Context, data models.Event) error
	DeleteEvent(ctx context.Context, query models.Event) error
	GetEvent(ctx context.Context, query models.Event) (*models.Event, error)
	GetEvents(ctx context.Context, query models.Event) ([]*models.Event, error)
	GetActiveEvents(ctx context.Context) ([]*models.Event, error)

	CreateOrUpdatePayment(ctx context.Context, data models.Payment) error
	DeletePayment(ctx context.Context, query models.Payment) error
	GetPayment(ctx context.Context, query models.Payment) (*models.Payment, error)
	GetPayments(ctx context.Context, query models.Payment) ([]*models.Payment, error)

	CreateOrUpdateBooking(ctx context.Context, data models.Booking) error
	DeleteBooking(ctx context.Context, query models.Booking) error
	GetBooking(ctx context.Context, query models.Booking) (*models.Booking, error)
	GetBookings(ctx context.Context, query models.Booking) ([]*models.Booking, error)
}

// New creates a new Database instance for repository
func New() (Repository, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}

	if err = migrateDB(db); err != nil {
		return nil, err
	}

	return &impl{db}, nil
}
