package main

import (
	"log"

	"github.com/zumosik/jwt-auth-golang/config"
	"github.com/zumosik/jwt-auth-golang/db"
	"github.com/zumosik/jwt-auth-golang/internal/user"
	"github.com/zumosik/jwt-auth-golang/router"
)

func main() {
	cfgPath, err := config.ParseFlags()
	if err != nil {
		log.Fatalf("Could not parse flags: %s", err)
	}
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatalf("Could not create config: %s", err)
	}

	dbConn, err := db.NewDatabase(cfg.Database.Url)
	if err != nil {
		log.Fatalf("Could not connect to db: %s", err)
	}

	uRep := user.NewRepostory(dbConn.GetDB())
	uSvc := user.NewService(uRep)
	uHandler := user.NewHandler(uSvc)

	router.InitRouter(uHandler)
	router.Start(cfg.Server.Port)
}
