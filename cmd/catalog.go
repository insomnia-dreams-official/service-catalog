package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/insomnia-dreams-official/service-catalog/internal/server"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Create and run server.
func main() {
	// Create database connection
	db, err := sql.Open("pgx", createDBConnString())
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	// Run database migrations
	dir, _ := os.Getwd()
	source := "file:" + path.Join(dir, "db/migrations")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		source,
		viper.GetString(`database.name`),
		driver,
	)
	if err != nil {
		log.Fatalf("failed to create new migration instance, reason: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to do migration(s), reason: %v", err)
	}

	// Create and start server
	s := server.New(db)
	s.Run()
}

func createDBConnString() string {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	return connStr
}
