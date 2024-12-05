package main

import (
	"context"
	"database/sql"
	"dms/config"
	"dms/modules/users"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func run() error {
	dbConfig, err := config.NewDatabaseConfig()
	if err != nil {
		return err
	}

	db, err := sql.Open("postgres", dbConfig.ConnectionURI)
	if err != nil {
		return err
	}

	ctx := context.Background()
	userRepo := users.NewUserRepository(db)

	allUsers, err := userRepo.ListAll(ctx)
	if err != nil {
		return err
	}

	for _, user := range allUsers {
		fmt.Printf("%v\n", user)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
