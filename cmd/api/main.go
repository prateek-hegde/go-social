package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/prateek-hegde/go-social/internal/consts"
	"github.com/prateek-hegde/go-social/internal/db"
	"github.com/prateek-hegde/go-social/internal/env"
	"github.com/prateek-hegde/go-social/internal/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", consts.DEFAULT_DB_ADDR),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 10),
			maxIdelConn:  env.GetInt("DB_MAX_IDLE_CONNS", 5),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIMEOUT", "30min"),
		},
	}
	db, err := db.New(cfg.addr, cfg.db.maxOpenConns, cfg.db.maxIdelConn, cfg.db.maxIdleTime)

	if err != nil {
		log.Panic("DB connection failed", err)
	}

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	log.Fatal(app.run())
}
