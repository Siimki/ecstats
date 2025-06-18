package dataclean

import (
	"ecstats/backend/config"
	"ecstats/backend/models"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"bufio"

//	"golang.org/x/tools/go/cfg"
)




type Group struct {
	Position int
	BibNumber int 
}


func ReadResultFromFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Cant read file: ",fileName , " Err:", err)
	}
	return data
}


func PrepareResultsData(matches [][]string,c *config.Config) (results []models.Result) {

	for _, match := range matches {
		var birthYear int
		position, err := strconv.Atoi(match[c.RegexGroups.Position+1])
		if err != nil {
			fmt.Println(err, "Cant convert in PrepareResultsData", position, "change positsion to 0")
			position = 0
		}

		bibNumber, err := strconv.Atoi(match[c.RegexGroups.BibNumber+1])
		if err != nil {
			fmt.Println(err, "Cant convert bibnumber in PrepareResultsData")
		}

		if c.RegexGroups.BirthYear > -1 {
			birthYear, err = strconv.Atoi(match[c.RegexGroups.BirthYear+1])
			if err != nil {
				fmt.Println(err, "Cant convert birthyear in PrepareResultsData")
			}
		}
		
		firstName, lastName, err := ExtractNames(match[c.RegexGroups.FullName+1])
		if err != nil {
			log.Fatal(err)
		}
		match[c.RegexGroups.Time+1] = strings.Replace(match[c.RegexGroups.Time+1], ",", ".", 1)

		lastNameCapitalized := Capitalize(lastName)
		fmt.Println(birthYear)
		fmt.Println(firstName)
		fmt.Println(lastName)
		fmt.Println(c.Race.RaceID)
		fmt.Println(position)
		fmt.Println(match[c.RegexGroups.Time+1])
		fmt.Println(bibNumber)
		results = append(results, models.Result{
			FirstName: firstName,
			LastName:  lastNameCapitalized,
			//BirthYear: birthYear,
			RaceId:    c.Race.RaceID,
			Position:  position,
			Time:      match[c.RegexGroups.Time+1],
			BibNumber: bibNumber,
		})
		fmt.Println()

	}

	return results
}


func ExtractNames(fullName string) (fname string, lname string, err error) {
	splittedFullName := strings.Split(fullName, " ")
	if len(splittedFullName) == 2 {
		return splittedFullName[0], splittedFullName[1], nil
	} else if len(splittedFullName) == 3 {
		return splittedFullName[0] + " " + splittedFullName[1], splittedFullName[2], nil
	} else if len(splittedFullName) == 4 {
		return splittedFullName[0] + " " + splittedFullName[1] + " " + splittedFullName[2], splittedFullName[3], nil
	}
	return "", "", fmt.Errorf("unexpected name format: %s. Len of name: %s", fullName, fname )
}


func ExtractNamesNew(fullName string) (fname string, lname string, err error) {
	splittedFullName := strings.Split(fullName, " ")
	if len(splittedFullName) == 2 {
		return splittedFullName[0], splittedFullName[1], nil
	} else if len(splittedFullName) == 3 {
		return splittedFullName[0] + " " + splittedFullName[1], splittedFullName[2], nil
	} else if len(splittedFullName) == 4 {
		return splittedFullName[0] + " " + splittedFullName[1] + " " + splittedFullName[2], splittedFullName[3], nil
	}
	return "", "", fmt.Errorf("unexpected name format. %s", fullName)
}


func isUpperCase(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func PrepareRiderData(matches [][]string,c *config.Config) (riders []models.Rider) {
	var problematicNames []string

	for _, match := range matches {
		var err error
		if c.RegexGroups.BirthYear > -1 {
			fmt.Println("c.RegexGroups.Birthyear =", c.RegexGroups.BirthYear)
			//year, err := strconv.Atoi(match[c.RegexGroups.BirthYear+1])
			if err != nil {
				fmt.Println("Error converting string to int at year conversion")
				return
			}
		}
		if string(match[c.RegexGroups.Gender+1]) != "M" {
			match[c.RegexGroups.Gender+1] = "F"
		} 

		//	var nationalityCode string
		if c.RegexGroups.Nationality > -1 {
			//nationalityCode = Capitalize(match[c.RegexGroups.Nationality+1])
		}

		problematicName := FindProblematicLine(match[c.RegexGroups.FullName+1])
		if problematicName != "" {
			problematicNames = append(problematicNames, problematicName)
		}
		firstName, lastName, err := ExtractNames(match[c.RegexGroups.FullName+1])
		if err != nil {
			log.Fatal(err)
		}
		lastNameCapitalized := Capitalize(lastName)
		//fmt.Println("Firstname: ", firstName, "Last:", lastNameCapitalized)
		riders = append(riders, models.Rider{
			LastName:    lastNameCapitalized,
			FirstName:   firstName,
			//BirthYear:   year,
			//Nationality: nationalityCode,
			Gender:      match[c.RegexGroups.Gender+1],
			// Team: 		match[c.RegexGroups.Team],
		})
		
	}

	fmt.Println(riders)
	if len(problematicNames) > 0 {
		fmt.Println(problematicNames)
		fmt.Println("Count of problematic names :", len(problematicNames))

		reader := bufio.NewReader(os.Stdin);
		fmt.Print("Continue anyway? (y/n): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input != "y" && input != "yes" {
			log.Fatal("stop program, solve problematic names")
		}
	
	}

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
	//fmt.Println("Fullname is:", fullName)
	if strings.Contains(fullName, "-") {
	//	fmt.Println(fullName, "contains hyphen")
		return fullName
	} else if splittedName := strings.Split(fullName, " "); len(splittedName) > 2 {
	//	fmt.Println(fullName, "contains longname")
		return fullName
	}
	return ""
}

func CheckDuplicates(matches [][]string) {

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

func CheckEdgeCases(c *config.Config,data string, matches [][]string) {
	var counter = 0
	for i, v := range matches {
		fmt.Println(c.RegexGroups.Position, " is pos ")
		fmt.Println(matches[i][c.RegexGroups.Position+1])
		if pos, err := strconv.Atoi(matches[i][c.RegexGroups.Position+1]); pos != counter {
			fmt.Println("Valuse of pos is:", pos, "valus of counter is:", counter)
			fmt.Println(err)
			fmt.Println("edgeCase problem", v)
			//log.Fatal()
		}
		fmt.Println("add counter")
		counter++
		fmt.Println(counter)
		if c.RegexGroups.BirthYear > -1 {
			fmt.Println("c.RegexGroups", c.RegexGroups.BirthYear)
			fmt.Println("Type of ")
			age, err := strconv.Atoi(matches[i][c.RegexGroups.BirthYear])
			if err != nil {
				log.Fatal("Come to checkEdgeCases and age check conversion")
			}
			AgeCheck(age)
		}

	}

	lines := strings.Split(string(data), "\n")
	fmt.Println(len(lines))
	fmt.Println(len(matches))

	if len(lines) == len(matches) {
		fmt.Println("All should be good. Now write return statement under this function")
	} else {
		log.Fatal("Lines and matches does not add up: Lines: ", len(lines), " Matches:", len(matches))
	}
	// Compile the regex
	re := regexp.MustCompile(c.RegexGroups.Regex)

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
		} else if len(rider.FirstName) < 2 {
			return fmt.Errorf("first name %s too short?", rider.FirstName)
		}

		if len(rider.LastName) > 50 {
			return fmt.Errorf("last name %s is too long", rider.LastName)
		} else if len(rider.LastName) < 2 {
			fmt.Println("Full name on short name is ", rider.FirstName, rider.LastName)
			return fmt.Errorf("last name %s too short?", rider.LastName)
		}
		
		if rider.BirthYear != 0 {
			if rider.BirthYear > 2026 || rider.BirthYear < 1901 {
				return fmt.Errorf("invalid birth year: %d", rider.BirthYear)
			}
		}

		if rider.Nationality != "" {
			if !regexp.MustCompile(`^[A-Z]{3}$`).MatchString(rider.Nationality) {
				return fmt.Errorf("invalid nationality: %s", rider.Nationality)
			}
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


