package dataclean

import (
	"database/sql"
	"ecstats/backend/db"
	"ecstats/backend/models"
	"ecstats/backend/config"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)
//GCparser
func CleanBosch(conn *sql.DB, c *config.Config) {
	regexGcBosch := `\d+\s+((?:[\p{L}\p{M}’'-]+\s+){1,3}[\p{L}\p{M}’'-]+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`
	fmt.Println("Clean bosch")
	data := ReadResultFromFile(c.Race.FileToRead)
	re := regexp.MustCompile(regexGcBosch)
	matches := re.FindAllStringSubmatch(string(data), -1)

	var counter int
	const numStages = 6
	boschStages := make([][][]string, numStages)
	var riders []models.Rider
	start:= time.Now()
	for _, match := range matches {

		counter++
		fullName := match[1]
		for stageIdx := 0; stageIdx < numStages; stageIdx++ {
			stageResult := match[stageIdx+2] // stage1 is match[4], stage2 is match[5], etc.
			boschStages[stageIdx] = append(boschStages[stageIdx], []string{stageResult, fullName})
		}
		firstName, LastName, err := ExtractNamesInBosch(fullName)
		if err != nil {
			log.Fatal("Acccident in clean bosch extract names! ", err)
		}
		lastNameCapitalized := Capitalize(LastName)
		firstNameCapitalized := CapitalizeFirstLetter(firstName)
		count := db.DuplicateCheck(conn, firstNameCapitalized, lastNameCapitalized)
		if count > 1 {
			fmt.Println("Duplicate found:", firstNameCapitalized, lastNameCapitalized)
		}		
		riders = append(riders, models.Rider{
			FirstName: firstNameCapitalized,
			LastName: lastNameCapitalized,
		})

	}
	// Now that we have all the data, sort each stage's results
	SortBosch(boschStages)
	
	var results []models.Result // Reset results for each stag	
	//for stageIdx, stage := range boschStages {
		raceId := c.Race.RaceID 
		for position , row := range boschStages[0] {
			//fullName := GCresult[2]
			firstName, LastName, err := ExtractNamesInBosch(row[1])
			if err != nil {
				log.Fatal("Acccident in clean bosch extract names! ")
			}
			lastNameCapitalized := Capitalize(LastName)
			firstNameCapitalized := CapitalizeFirstLetter(firstName)
			points, err := strconv.Atoi(row[0])
			if err != nil {
				log.Fatal("Conversion points error in clean bosch.")
			}
			if points == 0 {
				continue
			}
		
			results = append(results, models.Result{
				FirstName: firstNameCapitalized,
				LastName:  lastNameCapitalized,
				RaceId:    raceId,
				Position:  position+1,			
				Points:		points,
			})

		//}
		fmt.Println(results, c.Race.RaceID)
		db.AddResultsWithoutTimeToDb(conn, results, raceId)
		results = nil
	}
		//db.AddResultsWithoutTimeToDb(conn, results, c.Race.RaceID)

	fmt.Println("Stage was:", c.Race.StageNumber, "Id was", c.Race.RaceID)
	elapsed := time.Since(start)
	fmt.Printf("execution time: %s\n", elapsed)
	//riders2 := AddBoschGcRidersToDb(conn, boschStages[0]) //2016 done
	//db.AddRidersToDBwithNameOnly(conn, riders)
	
}

func AddBoschGcRidersToDb(conn *sql.DB ,gcResults [][]string) (riders []models.Rider)   {
	for _, gcResults := range gcResults{
		firstName, lastName, err := ExtractNamesInBosch(gcResults[1])
		if err != nil {
			fmt.Println(err)
		}
		lastNameCapitalized := Capitalize(lastName)
		firstNameCapitalized := CapitalizeFirstLetter(firstName)
		riders = append(riders, models.Rider{
			LastName:    lastNameCapitalized,
			FirstName:   firstNameCapitalized,
		})
		fmt.Println("Firstname, last name:",riders, "X")
	}
	return riders
}


func SortBosch(results [][][]string) () {
	for _, stageResults := range results {
		sortStageResults(stageResults)
	}
	return 
}

func sortStageResults(results [][]string) {
	sort.Slice(results, func(i, j int) bool {
		resultI, err1 := strconv.Atoi(results[i][0])
		resultJ, err2 := strconv.Atoi(results[j][0])
		if err1 != nil || err2 != nil {
			return results[i][0] > results[j][0]
		}
		return resultI > resultJ

	})
}

func ExtractNamesInBosch(fullName string) (fname string, lname string, err error) {
	splittedFullName := strings.Split(fullName, "	")
	if len(splittedFullName) == 2 {
		return splittedFullName[0], splittedFullName[1], nil
	} else if len(splittedFullName) == 3 {
		return splittedFullName[0] + " " + splittedFullName[1], splittedFullName[2], nil
	} else if len(splittedFullName) == 4 {
		return splittedFullName[0] + " " + splittedFullName[1] + " " + splittedFullName[2], splittedFullName[3], nil
	}
	return "", "", fmt.Errorf("unexpected name format. %s", fullName)
}

