package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)
func main() {
	// Sample todos
	myTodos := todos{
		{
			title:     "Learn Go",
			completed: false,
			createdAt: time.Now(),
		},
		{
			title:     "Write blog post",
			completed: true,
			createdAt: time.Now().Add(-48 * time.Hour),
			completedAt: func() *time.Time {
				t := time.Now().Add(-24 * time.Hour)
				return &t
			}(),
		},
	}

	// Print the todos
	myTodos.printManual()

	// Connect to the database
	connStr := "postgres://postgres:12345@localhost:5432/postgres"
	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
 	fmt.Println("Successfully connected to the database!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
