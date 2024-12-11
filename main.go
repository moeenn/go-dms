package main

import (
	"database/sql"
	"dms/config"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func run() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	_, err = sql.Open("postgres", cfg.Database.ConnectionURI)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
