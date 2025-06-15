package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/ritikkoul0/stock-rpc/server/utils"

	"github.com/ritikkoul0/stock-rpc/server/operations/overview"
)

var DB *sql.DB

func InitializeConnection(ctx context.Context, config *utils.AppConfig) error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open DB connection: %w", err)
	}
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}
	DB = db
	return nil
}

func Insertoverview(ctx context.Context, overview overview.Overview) error {
	// Convert struct into JSON first
	bytes, err := json.Marshal(overview)
	if err != nil {
		log.Printf("❌ Failed to marshal overview: %v", err)
		return fmt.Errorf("failed to marshal overview: %w", err)
	}

	// Prepare your insert
	query := `
		INSERT INTO Overview (overview)
		VALUES ($1)
		RETURNING id
	`

	var id int
	err = DB.QueryRowContext(ctx, query, bytes).Scan(&id)
	if err != nil {
		log.Printf("❌ Failed to insert overview: %v", err)
		return fmt.Errorf("failed to insert overview: %w", err)
	}

	log.Printf("✅ Overview successfully inserted with ID %d", id)
	return nil
}
