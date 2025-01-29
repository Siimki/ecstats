package models

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
	Status string
}

const RaceId = 36
const FileToRead = "../results/TARTU-RATTAMARATON/2024-TARTU-RATTAMARATON"