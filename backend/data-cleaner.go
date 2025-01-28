package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"ecstats/backend/models"
)

var patternForRidersInfo = `([A-ZÄÖÜÕŠŽ]+(?:[\s-][A-ZÄÖÜÕŠŽ]+)?)\s+([A-ZÄÖÜÕŠŽ][a-zäöüõšž]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž]+)*)\s+(\d{4})\s+([A-Z]{3})\s+\S+\s+([MN](?:\sjuunior|[-\d]+)?)`
var patternForTeamsInfo = `(?m)^(?:[^\t]*\t){8}([^\t]*)\t\d+\s*$`
var patternForResults = `(?m)(\d+)\t(\d+)\t([A-ZÄÖÜÕŠŽ]+(?:[\s-][A-ZÄÖÜÕŠŽ]+)?)\s+([A-ZÄÖÜÕŠŽ][a-zäöüõšž]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž]+)*)\s+(\d{4})\s+[A-Z]{3}\s+(\S+)\s+[MN]\s?(?:juunior|\d{2}-\d{2})\t\d+\t([A-ZÄÖÜÕŠŽa-zäöüõšž\d]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž\d-\/]+)*)?`
var regex2022results = `(\d+|\D{3})?\t(\d+)\t([A-ZÄÖÜÕŠŽ\s]+(?:[\s-][A-ZÄÖÜÕŠŽ]+)?)\s+([A-ZÄÖÜÕŠŽ][a-zäöüõšž]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž]+)*)\s+(\d{4})\s+([A-Z]{3})\s+(\d+:\d+:\d+)?`
var regex2022riders = `(\d+|\D{3})?\t(\d+)\t([A-ZÄÖÜÕŠŽ\s]+(?:[\s-][A-ZÄÖÜÕŠŽ]+)?)\s+([A-ZÄÖÜÕŠŽ][a-zäöüõšž]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž]+)*)\s+(\d{4})\s+([A-Z]{3})\s+(\d+:\d+:\d+)?`
var regex2019 = `(\d+|\D{3})?\t(\d+)\t([A-ZÄÖÜÕŠŽ\s]+(?:[\s-][A-ZÄÖÜÕŠŽ]+)?)\s+([A-ZÄÖÜÕŠŽ][a-zäöüõšž]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž]+)*)\s+(\d{4})\s+([A-Z]{3})\s+(\d+:\d+:\d+)?\t([MN\d]+)`

func readResultFromFile() []byte {
	data, err := os.ReadFile(models.FileToRead)
	if err != nil {
		log.Fatal("Cant read file")
	}
	return data
}

func PrepareResultsData() (results []models.Result)  {
	data := readResultFromFile()
	
	re := regexp.MustCompile(regex2019)
	matches := re.FindAllStringSubmatch(string(data), -1)

	for _, match := range matches{
		position, err := strconv.Atoi(match[1]) 
		if err != nil {
			fmt.Println(err, "Cant convert in PrepareResultsData",position, "change positsion to 0")
			position = 0 
		}

		bibNumber, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println(err, "Cant convert bibnumber in PrepareResultsData")
		}

		birthYear, err := strconv.Atoi(match[5])
		if err != nil {
			fmt.Println(err, "Cant convert birthyear in PrepareResultsData")
		}
		
		results = append(results, models.Result{
			FirstName: match[4],
			LastName: match[3],
			BirthYear: birthYear,
			RaceId: models.RaceId,
			Position: position,
			Time: match[7],
			BibNumber: bibNumber,
			//Status: "DNF",
		})
	}
	return results
}

func PrepareTeamsData() {
	data := readResultFromFile()

	re := regexp.MustCompile(regex2019)

	matches := re.FindAllStringSubmatch(string(data), -1)

	for _, match := range matches {
		fmt.Println(match[1])
	}
}

func PrepareRiderData()(riders []models.Rider) {

	data := readResultFromFile()

	// Compile the regex
	re := regexp.MustCompile(regex2019)
	
	matches := re.FindAllStringSubmatch(string(data), -1)
	checkEdgeCases(string(data), matches, regex2019)
	checkDuplicates(matches)

	for _, match := range matches {

		year, err := strconv.Atoi(match[5])
		if err != nil {
			fmt.Println("Error converting string to int at year conversion")
			return
		}

		if string(match[8][0]) != "M" && string(match[8][0]) != "N" {
			fmt.Println("Other gender?")
			return
		}

		if string(match[8][0]) == "N" {
			match[8] = "F"
		} else {
			match[8] = "M"
		}

		riders = append(riders, models.Rider{
			LastName: match[3],
			FirstName: match[4],
			BirthYear: year,
			Nationality: match[6],
			Gender: match[8],
		})

	}
	
	
	return riders
}

func checkDuplicates(matches [][]string) {
	
	nameMap := make(map[string]int)
	duplicatesFound := false
	var duplicates []string

	for _, match := range matches {
		fullNameWithYear := fmt.Sprintf("%s %s %s", match[1], match[2], match[3])
		nameMap[fullNameWithYear]++
	}

	for name, count := range nameMap{
		if count > 1 {
			fmt.Println(name)
			duplicates = append(duplicates, name)
			duplicatesFound = true
		}
	}

	if !duplicatesFound {
		fmt.Println("No duplicates found!")
	} else {
		log.Fatal("Duplicates found!", duplicates)
	}

}

func checkEdgeCases(data string, matches [][]string, regex string) {
	var counter = 1
	for i, v := range matches {
		if pos, err := strconv.Atoi(matches[i][1]); pos != counter {
			fmt.Println(err)
			fmt.Println(v, "that was problem")
			log.Fatal()
		}
		counter++
		age, err := strconv.Atoi(matches[i][5])
		if err != nil {
			log.Fatal("Come to checkEdgeCases and age check conversion")
		}
		ageCheck(age)
	}
	lines := strings.Split(string(data), "\n")
	fmt.Println(len(lines))
	fmt.Println(len(matches))

	if len(lines) == len(matches) {
		fmt.Println("All should be good. Now write return statement under this function")
	} else {
		log.Fatal("Lines and matches does not add up")
	}
	// Compile the regex
	re := regexp.MustCompile(regex)

	for _, line := range lines {
		line = strings.TrimSpace(line) // Clean up extra spaces
		if line == "" {
			continue
		}

		match := re.FindStringSubmatch(line)
		if match == nil {
			fmt.Println("No match for line:", line) // Log unmatched lines
			continue
		}
	}
	fmt.Println("If nothing was printed then all is good")
}

func ageCheck(birthYear int) {
	if birthYear > 2010 && birthYear < 1910 {
		fmt.Println("Error: rider born in", birthYear)
		log.Fatal("Bruv check the birth year here")
	}
}

