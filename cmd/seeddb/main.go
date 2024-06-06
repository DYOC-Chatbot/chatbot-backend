package main

import (
	"backend/internal/configs"
	"backend/internal/database"
)

func main() {
	cfg, err := configs.GetConfig()
	if err != nil {
		panic(err)
	}

	database.SetupDb(cfg.GetDatabaseConfig())

	database.PopulateDb()
}
