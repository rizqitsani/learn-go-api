package config

import (
	"database/sql"
	"fmt"
	"time"
)

func NewDatabase(config *Config) *sql.DB {
	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
