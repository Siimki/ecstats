package db

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"ecstats/backend/models"
	"log"
)

func ConnectToDB() *sql.DB {
	connStr := "user=postgres password=admin dbname=ecstats sslmode=disable port=5433"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AddRidersToDB(db *sql.DB ,riders []models.Rider) {
	var counter = 0
	for _, rider := range riders {
		RiderId :=  QueryRiderId(db ,rider.FirstName, rider.LastName, rider.BirthYear)
		if RiderId < 1 {
			fmt.Println("Adding new rider to the db:", rider.FirstName, rider.LastName, rider.BirthYear)
			counter++
		} else {
			continue
		}

		_, err := db.Exec(
			`INSERT INTO riders (first_name, last_name, birth_year, nationality, gender)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (first_name, last_name, birth_year) DO NOTHING;`,
			rider.FirstName, rider.LastName, rider.BirthYear, rider.Nationality, rider.Gender,
		)
		if err != nil {
			fmt.Printf("Error insterting rider %s %s: %v\n", rider.FirstName, rider.LastName, err)
		}
	}
	fmt.Println(counter,"Riders added to the DB!")
}

func AddTeamsToDB(db *sql.DB ,riders []models.Rider) {
	var counter = 0
	for _, rider := range riders {
		if rider.Team == "" {
			continue
		}
		
		teamId := QueryTeamId(db, rider.Team)
		if teamId > 0{
			continue 
		}

		_, err := db.Exec(`INSERT INTO teams (name) VALUES ($1) ON CONFLICT (name) DO NOTHING`, rider.Team)
		if err != nil {
			fmt.Printf("Error inserting teams %s: %v\n", rider.Team, err)
		} else {
			counter++
			fmt.Println("Added new team:", rider.Team)
		}
	}
	fmt.Println("Total new teams added: ", counter)
}

func AddRiderTeamRelations(db *sql.DB, riders []models.Rider, year int) {
	
	for _, rider := range riders {
		if rider.Team == "" {
			continue 
		}

		riderId := QueryRiderId(db, rider.FirstName, rider.LastName, rider.BirthYear)
		if riderId < 1 {
			fmt.Println("skipping rider(NOT FOUND)", rider.FirstName, rider.LastName)
			continue
		}

		teamId := QueryTeamId(db, rider.Team)
		if teamId < 1 {
			fmt.Println("skipping team(not found):", rider.Team)
			continue
		}

		_, err := db.Exec(`INSERT INTO rider_teams (rider_id, team_id, year) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`,
		riderId, teamId, year)
		if err != nil {
			fmt.Printf("Error inserting rider-team link for %s: %v\n", rider.FirstName, err)
		} else {
			//fmt.Printf("Linked %s %s to %s for %d\n", rider.FirstName, rider.LastName, rider.Team, year)
		}
	}

}


func AddResultsToDb(db *sql.DB, results []models.Result) {
	fmt.Println("Adding results to DB!")
	var counter = 0 
	for _, result := range results {

		RiderId :=  QueryRiderId(db ,result.FirstName, result.LastName, result.BirthYear)
		if RiderId < 1 {
			fmt.Printf("Rider %s %s (born %d) not found in the database. Cannot add result.\n", 
			result.FirstName, result.LastName, result.BirthYear)
			continue
		}
		if result.Position == 0 {
			timeValue := "00:00:00" // Represents DNF with a valid interval
			_, err := db.Exec (
				`INSERT INTO results (race_id,rider_id, position, time, race_number, status)
				VALUES ($1, $2, $3, $4, $5, $6)
				ON CONFLICT (race_id, rider_id) DO NOTHING;`,
				models.RaceId, RiderId, result.Position, timeValue, result.BibNumber,result.Status,
			)
			if err != nil {
				fmt.Printf("Error insterting result for  %s %s: %v\n and pos is 0", result.FirstName, result.LastName, err)
			}
			continue
		}
		_, err := db.Exec (
			`INSERT INTO results (race_id,rider_id, position, time, race_number)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (race_id, rider_id) DO NOTHING;`,
			models.RaceId, RiderId, result.Position, result.Time, result.BibNumber,
		)
		counter++
		if err != nil {
			fmt.Printf("Error insterting result for  %s %s: %v\n", result.FirstName, result.LastName, err)
		}
	}
	fmt.Println(counter ,"Results added to DB!")
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
			fmt.Printf("no rider found for %s %s, born in %v\n", firstName, lastName, birthYear)
			return 0
		}
		fmt.Printf("Error querying rider_id for %s %s: %v\n", firstName, lastName, err)
		return 0
	}
	return riderId
}

func QueryTeamId(db *sql.DB, teamName string) (teamId int) {
	err := db.QueryRow(
		`SELECT id FROM teams WHERE name = $1`,
		teamName,
	).Scan(&teamId)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No team found for %s: %v\n", teamName, err)
			return 0
		}
		fmt.Printf("Error querying team_id for %s: %v\n", teamName, err)
		return 0
	}

	//fmt.Printf("Found team ID %d for %s\n", teamId, teamName)
	return teamId
}

