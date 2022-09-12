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

	db, err := db.NewDatabase()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	// cmtService.PostComment(context.Background(), comment.Comment{
	// 	ID:     "71c5d074-b6cf-11ec-b909-0242ac120002",
	// 	Slug:   "test-slug",
	// 	Body:   "test-body",
	// 	Author: "test-author",
	// })

	fmt.Println(cmtService.GetComment(context.Background(), "35270f54-a188-49b9-9c6e-4d853e72eca6"))

	fmt.Println("successfully connected to database")

	return nil
}

func main() {
	fmt.Println("Hello, world!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
