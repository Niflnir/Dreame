package database

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "dreame"
)

var DBCon *sql.DB

func ConnectToDB() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected to database:%s!\n", dbname)

	// downMigrations := os.Args[1]
	// if downMigrations == "down" {
	//   goose.Down(db, "migrations")
	//   db.Close()
	//   os.Exit(1)
	// }

	// Set SQL dialect to postgresql
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	fmt.Println("Successfully changed dialect to postgres!")

	// Apply migrations to db
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	fmt.Println("Successfully migrated database!")

	return db
}
