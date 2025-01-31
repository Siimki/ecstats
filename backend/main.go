package main

import (
	"ecstats/backend/db"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"ecstats/backend/config"
	"ecstats/backend/dataclean"

)

func main() {

	conn := db.ConnectToDB()
	defer conn.Close()

	err := conn.Ping()
	if err != nil {
		log.Fatal("Database not reachable: ", err)
	}

	fmt.Println("Database connect succesfully!")
	
	dataclean.CapitalizePopularLastNames()
	//I added all riders so I am fine with that currently.

	riders := dataclean.PrepareRiderData()
	db.AddRidersToDB(conn, riders)
	
	results := dataclean.PrepareResultsData()
	db.AddResultsToDb(conn, results)

	db.AddTeamsToDB(conn, riders)

	db.AddRiderTeamRelations(conn, riders, config.Year)
	
	fmt.Println("Job finished!")
}


