package db

import (
	"L0/pkg/customLogger"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPsql(log customLogger.Logger) *sqlx.DB {
	psql, err := sqlx.Open("postgres", fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s sslmode=%s",
		"localhost", "postgres",
		"postgres", "password",
		"5432", "disable"))
	if err != nil {
		log.Print(err.Error())
	}

	if err := psql.Ping(); err != nil {
		log.Print(err.Error())
	}

	return psql
}
