package main

import (
	"context"
	"fmt"
	"go-rest-api/internal/comment"
	"go-rest-api/internal/db"
	transportHttp "go-rest-api/internal/transport/http"
)

func Run() error {
	fmt.Println("Starting our application")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to DB")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}
	fmt.Println("Successfully Pinged our DB")

	cmtService := comment.NewService(db)
	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"12345",
	))

	return nil
}
func main() {
	fmt.Println("\nGo Rest Api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
