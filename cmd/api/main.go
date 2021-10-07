package main

import (
	"projectName/internal/api"
	"projectName/pkg/config"
	"projectName/pkg/data"
)

func main() {
	cfg := config.New()
	db := data.NewMongoConnection(cfg)
	defer db.Disconnect()
	application := api.New(cfg, db.Client)
	application.Start()
}
