package main

import (
	"fmt"
	_ "github.com/lib/pq" // Import the pq package for PostgreSQL
	"log"
	"net/http"
)

func x2() {
	
	db := ConnectToDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Fatal("Database not reachable: ", err)
	}

	fmt.Println("Database connect succesfully!")

	fmt.Println("Starting server")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Welcome to the Ecstats!")
	// })

	fmt.Println("Server started	at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
