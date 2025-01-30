package models

type Rider struct {
	LastName   string
	FirstName  string
	BirthYear  int
	Nationality string
	Gender      string
	Team 		string
}

type Team struct {
	Name string
	Year int
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
	Status string
}


//last TRR was id 36
//42 was Rakke
const RaceId = 43
const FileToRead = "../results/BOSCH/output.txt"

var PatternForRidersInfo = `([A-ZÄÖÜÕŠŽ]+(?:[\s-][A-ZÄÖÜÕŠŽ]+)?)\s+([A-ZÄÖÜÕŠŽ][a-zäöüõšž]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž]+)*)\s+(\d{4})\s+([A-Z]{3})\s+\S+\s+([MN](?:\sjuunior|[-\d]+)?)`
var PatternForTeamsInfo = `(?m)^(?:[^\t]*\t){8}([^\t]*)\t\d+\s*$`
var PatternForResults = `(?m)(\d+)\t(\d+)\t([A-ZÄÖÜÕŠŽ]+(?:[\s-][A-ZÄÖÜÕŠŽ]+)?)\s+([A-ZÄÖÜÕŠŽ][a-zäöüõšž]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž]+)*)\s+(\d{4})\s+[A-Z]{3}\s+(\S+)\s+[MN]\s?(?:juunior|\d{2}-\d{2})\t\d+\t([A-ZÄÖÜÕŠŽa-zäöüõšž\d]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž\d-\/]+)*)?`
var Pegex2022results = `(\d+|\D{3})?\t(\d+)\t([A-ZÄÖÜÕŠŽ\s]+(?:[\s-][A-ZÄÖÜÕŠŽ]+)?)\s+([A-ZÄÖÜÕŠŽ][a-zäöüõšž]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž]+)*)\s+(\d{4})\s+([A-Z]{3})\s+(\d+:\d+:\d+)?`
var Pegex2022riders = `(\d+|\D{3})?\t(\d+)\t([A-ZÄÖÜÕŠŽ\s]+(?:[\s-][A-ZÄÖÜÕŠŽ]+)?)\s+([A-ZÄÖÜÕŠŽ][a-zäöüõšž]+(?:[\s-][A-ZÄÖÜÕŠŽa-zäöüõšž]+)*)\s+(\d{4})\s+([A-Z]{3})\s+(\d+:\d+:\d+)?`