package main

import (
	"database/sql"
	"fmt"
	"github.com/afarid/todo/api"
	database "github.com/afarid/todo/db/sqlc"
	"github.com/afarid/todo/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	var dbConnectionUrl string
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	if config.DBDriver == "postgres" {
		dbConnectionUrl = fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable",
			config.DBUser,
			config.DBPassword,
			config.DBHost,
			config.DBName,
		)

	} else {
		log.Fatal("Unsupported database driver", config.DBDriver)
	}
	conn, err := sql.Open(config.DBDriver, dbConnectionUrl)
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	runDBMigration(config.MigrationURL, dbConnectionUrl)

	dbClient := database.New(conn)
	server := api.NewServer(dbClient)
	server.Run()
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated successfully")
}
