package main

import (
	"fmt"

	"github.com/MrBooi/go-rest-api/internal/comment"
	"github.com/MrBooi/go-rest-api/internal/db"
	transportHttp "github.com/MrBooi/go-rest-api/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)

	if err := httpHandler.Serve(); err != nil {
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
