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
)

func CleanBosch(conn *sql.DB) {
	//group 1 is gc place
	//group 2 is names
	//group 3 is gc points
	// 4-11 are stages
	//namesToTest := []string{"Mari-Liis Mõttus", "Mari Mõttus-Ottender", "Mari-Liis Mõttus-Ottender", "Mari Liis Mõttus", "Mari Liis Mõttus-Ottender", "Mathieu van der Poel"}
	regexGcBosch := `(\d+).\s+([\w\-\ \\-äõüöšŠĀī]+)\t(\d+)\t(\d+)\t(\d+)\t(\d+)\t(\d+)\t(\d+)\t(\d+)\t(\d+)\t(\d+)\t(\d+)\t(\d+)\n`
	fmt.Println("Clean bosch")
	data := ReadResultFromFile()
	re := regexp.MustCompile(regexGcBosch)
	matches := re.FindAllStringSubmatch(string(data), -1)

	var counter int
	const numStages = 9
	var results []models.Result
	boschStages := make([][][]string, numStages)

	for _, match := range matches {

		counter++
		fullName := match[2]
	//	fmt.Println(match[0])
		for stageIdx := 0; stageIdx < numStages; stageIdx++ {
			stageResult := match[stageIdx+3] // stage1 is match[4], stage2 is match[5], etc.
			boschStages[stageIdx] = append(boschStages[stageIdx], []string{stageResult, fullName})
		}
		firstName, LastName, err := ExtractNamesInBosch(fullName)
		if err != nil {
			log.Fatal("Acccident in clean bosch extract names! ", err)
		}
		lastNameCapitalized := Capitalize(LastName)
		count := db.DuplicateCheck(conn, firstName, lastNameCapitalized)
		if count > 1 {
			fmt.Println("Duplicate found:", firstName, lastNameCapitalized)
		}		

	}
	// Now that we have all the data, sort each stage's results.
//	SortBosch(boschStages)
//	fmt.Println(boschStages)

	// for i := 1; i <= 8; i++ {

	// }

	for position , row := range boschStages[0] {
		//fullName := GCresult[2]
		firstName, LastName, err := ExtractNamesInBosch(row[1])
		if err != nil {
			log.Fatal("Acccident in clean bosch extract names! ")
		}
		lastNameCapitalized := Capitalize(LastName)
		points, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatal("Conversion points error in clean bosch.")
		}
		if points == 0 {
			continue
		}
		results = append(results, models.Result{
			FirstName: firstName,
			LastName:  lastNameCapitalized,
			RaceId:    config.RaceId,
			Position:  position+1,			
			Points:		points,
		})
	}
	// fmt.Println(results)
	// fmt.Println(results[0].FirstName)
	// fmt.Println(results[0].LastName)
	// fmt.Println(results[0].RaceId)
	// fmt.Println(results[0].Position)

	//db.AddResultsWithoutTimeToDb(conn, results, 78)

//	fmt.Println(len(boschStages))
	//AddBoschGcRidersToDb(conn, boschStages[0]) //2016 done
	

}

func AddBoschGcRidersToDb(conn *sql.DB ,gcResults [][]string)   {
	var riders []models.Rider
	for _, gcResults := range gcResults{
		firstName, lastName, err := ExtractNamesInBosch(gcResults[1])
		if err != nil {
			fmt.Println(err)
		}
		lastNameCapitalized := Capitalize(lastName)
		riders = append(riders, models.Rider{
			LastName:    lastNameCapitalized,
			FirstName:   firstName,
		})
		fmt.Println("Firstname, last name:",riders, "X")
	
	}
	db.AddRidersToDBwithNameOnly(conn, riders)
	
}



func SortBosch(results [][][]string) () {
	for _, stageResults := range results {
		sortStageResults(stageResults)
		//fmt.Printf("Sorted Stage %d Results: %v\n", i+1, stageResults)
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
	splittedFullName := strings.Split(fullName, " ")
	if len(splittedFullName) == 2 {
		return splittedFullName[0], splittedFullName[1], nil
	} else if len(splittedFullName) == 3 {
		return splittedFullName[0] + " " + splittedFullName[1], splittedFullName[2], nil
	} else if len(splittedFullName) == 4 {
		return splittedFullName[0] + " " + splittedFullName[1] + " " + splittedFullName[2], splittedFullName[3], nil
	}
	return "", "", fmt.Errorf("Unexpected name format. %s", fullName)
}
