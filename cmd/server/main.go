package main

import (
	"context"
	"fmt"

	"github.com/MrBooi/go-rest-api/internal/comment"
	"github.com/MrBooi/go-rest-api/internal/db"
)

// Run - is going to be executed when you run the program
func Run() error {
	fmt.Println("Starting up our application...")

	store, err := db.NewDatabase()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}

	if err := db.MigrateDB(store); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(store)
	fmt.Println(cmtService.GetComment(context.Background(), "71c5d074-b6cf-11ec-b909-0242ac120002"))

	fmt.Println("successfully connected to database")

	return nil
}

func main() {
	fmt.Println("Hello, world!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
