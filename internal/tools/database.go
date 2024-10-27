package tools

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

type PostgresDb struct {
	db *sql.DB
}

func NewDatabase() (*PostgresDb, error) {
	fmt.Println("Opening DB connection...")
	connectionString := "postgres://postgres:admin@localhost:5400/dev?sslmode=disable"

	postgresDb, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err = postgresDb.Ping(); err != nil {
		log.Fatal(err)
	}

	return &PostgresDb{db: postgresDb}, nil
}

func (d *PostgresDb) CloseDatabase() error {
	if d.db != nil {
		err := d.db.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
			return err
		}
		fmt.Println("Database connection closed successfully.")
	}
	return nil
}
