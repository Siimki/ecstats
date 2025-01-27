package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rider struct {
	LastName   string
	FirstName  string
	BirthYear  int
	Nationality string
	Gender      string
}

type Team struct {
	Name string
}

type Result struct {
	FirstName string
	LastName string
	BirthYear int
	RaceId int
	RiderId int
	Position int
	Time string
	BibNumber int
}

var patternForRidersInfo = `([A-ZÄÖÜÕŠŽ]+(?:[\s-][A-ZÄÖÜÕŠŽ]+)?)\s+([A-ZÄÖÜÕŠŽ][a-zäöüõšž]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž]+)*)\s+(\d{4})\s+([A-Z]{3})\s+\S+\s+([MN](?:\sjuunior|[-\d]+)?)`
var patternForTeamsInfo = `(?m)^(?:[^\t]*\t){8}([^\t]*)\t\d+\s*$`
var patternForResults = `(?m)(\d+)\t(\d+)\t([A-ZÄÖÜÕŠŽ]+(?:[\s-][A-ZÄÖÜÕŠŽ]+)?)\s+([A-ZÄÖÜÕŠŽ][a-zäöüõšž]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž]+)*)\s+(\d{4})\s+[A-Z]{3}\s+(\S+)\s+[MN]\s?(?:juunior|\d{2}-\d{2})\t\d+\t([A-ZÄÖÜÕŠŽa-zäöüõšž\d]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž\d-\/]+)*)?`

func readResultFromFile() []byte {
	data, err := os.ReadFile("../results/2024-TARTU-RATTARALLI")
	if err != nil {
		fmt.Println("Cant read file")
	}
	return data
}

func PrepareResultsData() (results []Result)  {
	data := readResultFromFile()
	
	re := regexp.MustCompile(patternForResults)
	matches := re.FindAllStringSubmatch(string(data), -1)

	for _, match := range matches{
		position, err := strconv.Atoi(match[1]) 
		if err != nil {
			fmt.Println(err, "Cant convert in PrepareResultsData")
		}

		bibNumber, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println(err, "Cant convert in PrepareResultsData")
		}

		birthYear, err := strconv.Atoi(match[5])
		if err != nil {
			fmt.Println(err, "Cant convert in PrepareResultsData")
		}
		
		results = append(results, Result{
			FirstName: match[4],
			LastName: match[3],
			BirthYear: birthYear,
			RaceId: 3,
			Position: position,
			Time: match[6],
			BibNumber: bibNumber,
		})
	}
	return results
}

func PrepareTeamsData() {
	data := readResultFromFile()

	re := regexp.MustCompile(patternForTeamsInfo)

	matches := re.FindAllStringSubmatch(string(data), -1)

	for _, match := range matches {
		fmt.Println(match[1])
	}
}

func PrepareRiderData()(riders []Rider) {

	data := readResultFromFile()

	// Compile the regex
	re := regexp.MustCompile(patternForRidersInfo)
	
	matches := re.FindAllStringSubmatch(string(data), -1)

	for _, match := range matches {

		year, err := strconv.Atoi(match[3])
		if err != nil {
			fmt.Println("Error converting string to int at year conversion")
			return
		}

		if string(match[5][0]) != "M" && string(match[5][0]) != "N" {
			fmt.Println("Other gender?")
			return
		}

		if string(match[5][0]) == "N" {
			match[5] = "F"
		} else {
			match[5] = "M"
		}

		riders = append(riders, Rider{
			LastName: match[1],
			FirstName: match[2],
			BirthYear: year,
			Nationality: match[4],
			Gender: match[5],
		})

	}

	for _, rider := range riders{

		fmt.Printf("Last Name: %s, First Name: %s, Birth Year: %d, Nationality: %s, Gender: %s\n",
			rider.LastName, rider.FirstName, rider.BirthYear, rider.Nationality, rider.Gender)
	}
	
	checkEdgeCases(string(data), matches, patternForRidersInfo)
	checkDuplicates(matches)
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
		fmt.Println("Duplicates found!", duplicates)
	}

}

func checkEdgeCases(data string, matches [][]string, regex string) {
	lines := strings.Split(string(data), "\n")
	fmt.Println(len(lines))
	fmt.Println(len(matches))
	
	if len(lines) == len(matches) {
		fmt.Println("All should be good. Now write return statement under this function")
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

