package main

import (
	"hanifu.id/hansputera-factory/garudacbt-backend/config"
	"hanifu.id/hansputera-factory/garudacbt-backend/database"
	"hanifu.id/hansputera-factory/garudacbt-backend/server"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewGarudaDatabase(cfg)

	server.NewFiberServer(cfg, db).Start()
}
