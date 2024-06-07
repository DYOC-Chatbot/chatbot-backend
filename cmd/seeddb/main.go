package main

import (
	"backend/internal/configs"
	"backend/internal/database"
	"flag"
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	// Delete current database
	cfg, err := configs.GetConfig()
	if err != nil {
		panic(err)
	}

	database.DropDb(cfg.GetDatabaseConfig())

	// Import schema from migration
	var countFlag = flag.Int("count", 0, "number of database migrations to run. omit (or 0) for no limit")

	db, err := database.GetDb().DB()
	if err != nil {
		log.Fatal(err)
	}

	migrate.ExecMax(db, "postgres", &migrate.FileMigrationSource{
		Dir: "./migrations",
	}, migrate.Up, *countFlag)

	// Populate schema with data
	database.PopulateDb()
}
