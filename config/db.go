package config

import (
	"log"
	"os"

	"github.com/ggowtham2745/users_api/controllers/user"
	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "enterprisedb",
		Password: "cteladmin",
		Addr:     "10.10.10.164:5444",
		Database: "Central5116",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")

	user.InitiateDB(db)

	return db
}
