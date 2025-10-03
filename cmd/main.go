package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alroymuhammad/gomanager_be/internal/handlers"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=gomanagergo sslmode=disable password=postgres host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handlers.HelloHandler)
	fmt.Println("Server running in 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
