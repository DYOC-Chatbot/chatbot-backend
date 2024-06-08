package database

import (
	"backend/internal/configs"
	"backend/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var globalDb *gorm.DB

func SetupDb(cfg *configs.PostgresConfig) {
	dsn, err := BuildDsn(cfg)
	if err != nil {
		panic("Error building the DSN.")
	}

	gormCfg := GetConfig()
	db, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		panic("Error opening the database.")
	}

	globalDb = db
}

// Assumption: SetupDb is called before this function
func GetDb() *gorm.DB {
	return globalDb.Session(&gorm.Session{NewDB: true})
}

func DropDb(cfg *configs.PostgresConfig) {
	// Modify the DSN to connect to the 'postgres' database
	var originalDb = cfg.PostgresDb

	cfg.PostgresDb = "postgres"
	dsn, err := BuildDsn(cfg)
	if err != nil {
		log.Fatalf("Error building the DSN: %v", err)
	}

	// Connect to the 'postgres' database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Terminate all ongoing processes
	terminateConnectionsQuery := fmt.Sprintf(`
		SELECT pg_terminate_backend(pg_stat_activity.pid)
		FROM pg_stat_activity
		WHERE pg_stat_activity.datname = '%s' AND pid <> pg_backend_pid();
	`, originalDb)
	if err := db.Exec(terminateConnectionsQuery).Error; err != nil {
		log.Fatalf("Error terminating connections: %v", err)
	}

	// Drop the target database
	dropDbQuery := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", originalDb)
	if err := db.Exec(dropDbQuery).Error; err != nil {
		log.Fatalf("Error dropping the database: %v", err)
	}

	// Create the target database
	createDbQuery := fmt.Sprintf("CREATE DATABASE %s;", originalDb)
	if err := db.Exec(createDbQuery).Error; err != nil {
		log.Fatalf("Error creating the database: %v", err)
	}

	// Close the connection to the 'postgres' database
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting raw DB connection: %v", err)
	}
	dbSQL.Close()

	// Update DSN to connect to the target database
	cfg.PostgresDb = originalDb

	// Setup connection to db
	SetupDb(cfg)
}

func PopulateDb() {
	// clears all values from the table and populates sample data
	model.PopulateUsers(globalDb)
	model.PopulateChats(globalDb)
	model.PopulateBookings(globalDb)
	model.PopulateRequestQueries(globalDb)
	model.PopulateMessages(globalDb)
}
