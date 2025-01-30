package main

import (
	// "database/sql"
	// "ecstats/backend/models"
	"ecstats/backend/dataclean"
	"ecstats/backend/db"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	conn := db.ConnectToDB()
	defer conn.Close()

	err := conn.Ping()
	if err != nil {
		log.Fatal("Database not reachable: ", err)
	}

	fmt.Println("Database connect succesfully!")
	
	//I added all riders so I am fine with that currently. 



	riders := dataclean.PrepareRiderData()
	db.AddRidersToDB(conn, riders)
	
	results := dataclean.PrepareResultsData()
	db.AddResultsToDb(conn, results)

	db.AddTeamsToDB(conn, riders)

	db.AddRiderTeamRelations(conn, riders, 2022)
	
	fmt.Println("Job finished!")
}


