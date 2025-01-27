package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
)

func main() {
	db := ConnectToDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Fatal("Database not reachable: ", err)
	}

	fmt.Println("Database connect succesfully!")
	
	//I added all riders so I am fine with that currently. 

	riders := PrepareRiderData()
	AddRidersToDB(db, riders)

	results := PrepareResultsData()
	AddResultsToDb(db, results)
}

func AddRidersToDB(db *sql.DB ,riders []Rider) {
	for _, rider := range riders {
		
		_, err := db.Exec(
			`INSERT INTO riders (first_name, last_name, birth_year, nationality, gender)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT DO NOTHING;`,
			rider.FirstName, rider.LastName, rider.BirthYear, rider.Nationality, rider.Gender,
		)
		if err != nil {
			fmt.Printf("Error insterting rider %s %s: %v\n", rider.FirstName, rider.LastName, err)
		}
	}
	fmt.Println("Riders added to the DB!")
}

func AddResultsToDb(db *sql.DB, results []Result) {

	for _, result := range results {

		RiderId :=  QueryRiderId(db ,result.FirstName, result.LastName, result.BirthYear)
		_, err := db.Exec (
			`INSERT INTO results (race_id,rider_id, position, time, race_number)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT DO NOTHING;`,
			3, RiderId, result.Position, result.Time, result.BibNumber,
		)
		if err != nil {
			fmt.Printf("Error insterting result for  %s %s: %v\n", result.FirstName, result.LastName, err)
		}
	}
	// fmt.Println(results[7].BibNumber)
	// fmt.Println(results[7].FirstName)
	// fmt.Println(results[7].LastName)
	// fmt.Println("Len of results", len(results))
}

func QueryRiderId(db *sql.DB,firstName string, lastName string, birthYear int) (riderId int){
	
	err := db.QueryRow(
		`SELECT id FROM riders WHERE first_name = $1 AND last_name = $2 AND birth_year = $3`,
		firstName, lastName, birthYear,
	).Scan(&riderId)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Birthyear", birthYear)
			fmt.Println("First name", firstName, "lastname" , lastName)
			fmt.Printf("no rider found for %s %s, born in %s\n", firstName, lastName, birthYear)
			return 0
		}
		fmt.Printf("Error querying rider_id for %s %s: %v\n", firstName, lastName, err)
		return 0
	}
	return riderId
}


func ConnectToDB() *sql.DB {
	connStr := "user=postgres password=admin dbname=ecstats sslmode=disable port=5433"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InsertRidersToDB() {

}
