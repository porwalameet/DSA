package main

import (
	"cleancode/config"
	"cleancode/database"
	"cleancode/server"
	"log"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	if err := db.CheckConnection(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	server.NewEchoServer(conf, db).Start()
}
