package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (sp *serviceProvider) GetDB() *sqlx.DB {
	var err error
	if sp.db == nil {
		dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", sp.cfg.DatabaseUsername, sp.cfg.DatabasePassword, sp.cfg.DatabaseHost, sp.cfg.DatabasePort, sp.cfg.DatabaseName)
		sp.db, err = sqlx.Open("postgres", dsn)
		if err != nil {
			panic(err)
		}

	}
	return sp.db
}
