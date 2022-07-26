package main

import (
	"fmt"
	"github.com/TutorialEdge/go-rest-api-course/internal/db"
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
	return nil
}
func main() {
	fmt.Println("\nGo Rest Api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
