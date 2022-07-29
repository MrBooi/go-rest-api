package main

import (
	"fmt"

	"github.com/MrBooi/go-rest-api/internal/db"
)

// Run - is going to be executed when you run the program
func Run() error {
	fmt.Println("Starting up our application...")

	database, err := db.NewDatabase()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}

	if err := db.MigrateDB(database); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully connected to database")

	return nil
}

func main() {
	fmt.Println("Hello, world!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
