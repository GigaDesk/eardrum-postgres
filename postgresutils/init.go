package postgresutils

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresInstance struct {
	Db *gorm.DB
}

func (p *PostgresInstance) Init(url string) error {
	// Get the database URL from the environment variable provided by Render
	databaseURL := url
	if databaseURL == "" {
		return errors.New("DATABASE_URL environment variable not set.")
	}

	// Establish the database connection using the URL
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return err
	}

	p.Db = db
	fmt.Println("Successfully connected to PostgreSQL database via Render URL!")
	return nil

	// You can now use the 'db' variable to interact with your database using GORM
	// ... your database operations ...
}
