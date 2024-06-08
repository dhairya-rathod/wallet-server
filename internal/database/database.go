package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"

	config "wallet-server/internal/config"
)

type Service struct {
	db *sql.DB
}

var (
	dbConfig   = config.LoadConfig().DB
	dbInstance *Service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return *dbInstance
	}
	connStr := dbConfig.ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &Service{
		db: db,
	}

	return *dbInstance
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s Service) Close() error {
	log.Printf("Disconnected from database: %s", dbConfig.DBName)
	return s.db.Close()
}
