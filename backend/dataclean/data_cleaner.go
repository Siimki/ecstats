package dataclean

import (
	"ecstats/backend/models"
	"ecstats/backend/config"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)


var regexBosch = "ADD REGEX HERE"

func ReadResultFromFile() []byte {
	data, err := os.ReadFile(config.FileToRead)
	if err != nil {
		log.Fatal("Cant read file")
	}
	return data
}

func PrepareResultsData() (results []models.Result) {

	data := ReadResultFromFile()
	re := regexp.MustCompile(regexBosch)
	matches := re.FindAllStringSubmatch(string(data), -1)

	for _, match := range matches {

		position, err := strconv.Atoi(match[8])
		if err != nil {
			fmt.Println(err, "Cant convert in PrepareResultsData", position, "change positsion to 0")
			position = 0
		}

		bibNumber, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println(err, "Cant convert bibnumber in PrepareResultsData")
		}

		birthYear, err := strconv.Atoi(match[4])
		if err != nil {
			fmt.Println(err, "Cant convert birthyear in PrepareResultsData")
		}

		firstName, lastName := ExtractNames(match[2])
		lastNameCapitalized := Capitalize(lastName)

		results = append(results, models.Result{
			FirstName: firstName,
			LastName:  lastNameCapitalized,
			BirthYear: birthYear,
			RaceId:    config.RaceId,
			Position:  position,
			Time:      match[7],
			BibNumber: bibNumber,
		})
		
	}
	return results
}

func ExtractNames(fullName string) (string, string) {
	
	splittedFullName := strings.Split(fullName, " ")
	if len(splittedFullName) == 2 {
		return splittedFullName[0], splittedFullName[1]
	}
	words := strings.Fields(fullName)
	var firstNameParts []string
	var lastNameParts []string
	foundLastName := false

	for _, word := range words {
		if isUpperCase(word) {
			foundLastName = true
		}
		if foundLastName {
			lastNameParts = append(lastNameParts, word)
		} else {
			firstNameParts = append(firstNameParts, word)
		}
	}

	firstName := strings.Join(firstNameParts, " ")
	lastName := strings.Join(lastNameParts, " ")

	return firstName, lastName
}

func isUpperCase(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func PrepareRiderData() (riders []models.Rider) {
	data := ReadResultFromFile()
    var problematicNames []string
	re := regexp.MustCompile(regexBosch)

	matches := re.FindAllStringSubmatch(string(data), -1)
	checkEdgeCases(string(data), matches, regexBosch)
	checkDuplicates(matches)

	for _, match := range matches {

		year, err := strconv.Atoi(match[4])
		if err != nil {
			fmt.Println("Error converting string to int at year conversion")
			return
		}

		if string(match[3]) != "Mees" && string(match[3]) != "Naine" {
			fmt.Println("Other gender?")
			return
		}

		if string(match[3]) == "Mees" {
			match[3] = "M"
		} else {
			match[3] = "F"
		}
		nationalityCode := Capitalize(match[5])
		problematicName := FindProblematicLine(match[2])
		if problematicName != "" {
			problematicNames = append(problematicNames, problematicName)
		}
		firstName, lastName := ExtractNames(match[2])
		lastNameCapitalized := Capitalize(lastName)
		riders = append(riders, models.Rider{
			LastName:    lastNameCapitalized,
			FirstName:   firstName,
			BirthYear:   year,
			Nationality: nationalityCode,
			Gender:      match[3],
			Team: 		match[6],
		})
		
	}
	fmt.Println(problematicNames)
	fmt.Println("Count of problematic names :", len(problematicNames))
	//log.Fatal("stop program, solve problematic names")


	err := ValidateRider(riders)
	if err != nil {
		log.Fatal("Failed to validate riders", err)
	}

	return riders
}



func Capitalize(s string) (nationalityCode string){
	nationalityCode = strings.ToUpper(s)
	return nationalityCode
}

func FindProblematicLine(fullName string)  (probleMaticName string ) {

	if strings.Contains(fullName, "-") {
	//	fmt.Println(fullName, "contains hyphen")
		return fullName
	} else if splittedName := strings.Split(fullName, " "); len(splittedName) > 2 {
	//	fmt.Println(fullName, "contains longname")
		return fullName
	}
	return ""
}

func checkDuplicates(matches [][]string) {

	nameMap := make(map[string]int)
	duplicatesFound := false
	var duplicates []string

	for _, match := range matches {
		fullNameWithYear := fmt.Sprintf("%s %s %s", match[1], match[2], match[3])
		nameMap[fullNameWithYear]++
	}

	for name, count := range nameMap {
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
		// fmt.Println(v[1])
		// fmt.Println(v[2])
		// fmt.Println(v[3])
		// fmt.Println(v[4])
		// fmt.Println(v[5])
		// fmt.Println(v[6])
		// fmt.Println(v[7])
		// fmt.Println(v[8])
		// fmt.Println("value of I is ")

		//m[i][8] is last element and in bosch it is positision
		//TRR is vice-versa
		if pos, err := strconv.Atoi(matches[i][8]); pos != counter {
			fmt.Println("Valuse of pos is:", pos, "valus of counter is:", counter)
			fmt.Println(err)
			fmt.Println("edgeCase problem", v)
			log.Fatal()
		}
		counter++
		age, err := strconv.Atoi(matches[i][4])
		if err != nil {
			log.Fatal("Come to checkEdgeCases and age check conversion")
		}
		AgeCheck(age)
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

func AgeCheck(birthYear int) {
	if birthYear > 2010 && birthYear < 1910 {
		fmt.Println("Error: rider born in", birthYear)
		log.Fatal("Bruv check the birth year here")
	}
}

func CapitalizeFirstLetter(name string) string {
	if name == "" {
		return name
	}

	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])

	for i := 1; i < len(runes); i++ {
		if runes[i-1] == ' ' || runes[i-1] == '-' {
			runes[i] = unicode.ToUpper(runes[i])
		} else {
			runes[i] = unicode.ToLower(runes[i])
		}
	}

	return string(runes)
}


func ValidateRider(riders []models.Rider) error {

	for _, rider := range riders {
		if len(rider.FirstName) > 50 {
			return fmt.Errorf("first name %s is too long", rider.FirstName)
		}

		if len(rider.LastName) > 50 {
			return fmt.Errorf("last name %s is too long", rider.LastName)
		}

		if rider.BirthYear > 2026 || rider.BirthYear < 1901 {
			return fmt.Errorf("invalid birth year: %d", rider.BirthYear)
		}

		if !regexp.MustCompile(`^[A-Z]{3}$`).MatchString(rider.Nationality) {
			return fmt.Errorf("invalid nationality: %s", rider.Nationality)
		}

		validGenders := map[string]bool{"M": true, "F": true, "O": true, "": true}
		if !validGenders[rider.Gender] {
			return fmt.Errorf("invalid gender: %s", rider.Gender)
		}

			// Check if last name is FULLY uppercase
		if !isUpperCase(rider.LastName) {
			return fmt.Errorf("last name %s must be fully uppercase", rider.LastName)
		}


	}

	return nil
}


