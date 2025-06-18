package main

import (
	"ecstats/backend/db"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"ecstats/backend/config"
	"ecstats/backend/dataclean"
	"regexp"

)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	conn := db.ConnectToDB(cfg)
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Fatal("Database not reachable: ", err)
	}

	fmt.Println("Database connect succesfully!")
	
	data := dataclean.ReadResultFromFile(cfg.Race.FileToRead)
	re := regexp.MustCompile(cfg.RegexGroups.Regex)

	matches := re.FindAllStringSubmatch(string(data), -1)

	//Remove check edge cases for TTP(Pair Time Trial)

	dataclean.CheckEdgeCases(cfg ,string(data), matches)
	dataclean.CheckDuplicates(matches)
	riders := dataclean.PrepareRiderData(matches, cfg)

	err = dataclean.ValidateRider(riders)
	if err != nil {
		log.Fatal("Failed to validate riders", err)
	}

	if cfg.RegexGroups.BirthYear > -1 && cfg.RegexGroups.Nationality > -1 {
		db.AddRidersToDB(conn, riders)
	} else {
		db.AddRidersToDBwithNameOnly(conn, riders)
	}

	//make if-else statements when to use bosch and when prepare. 

	//dataclean.CleanBosch(conn, cfg)
	//dataclean.CapitalizePopularLastNames(cfg)

	// I added all riders so I am fine with that currently.

	 results := dataclean.PrepareResultsData(matches, cfg)
	 fmt.Println(results)
	 db.AddResultsToDb(conn, results, cfg)

	// To add with teams
	// db.AddTeamsToDB(conn, riders)
	// db.AddRiderTeamRelations(conn, riders, config.Year)
	fmt.Println("Job finished!")

} 







