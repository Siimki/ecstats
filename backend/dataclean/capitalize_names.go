package dataclean

import (
	"fmt"
	"io/ioutil"
	"strings"
	"ecstats/backend/config"
	"log"
)

func CapitalizePopularLastNames(c *config.Config) {

	content, err := ioutil.ReadFile(c.Race.FileToRead)
	if err != nil {
		log.Fatal("Cant read file in capitalize")
		return
	}

	text := string(content)
	for _, name := range config.PopularNamesToCapitalize {
		text = strings.ReplaceAll(text, name, strings.ToUpper(name))
	}

	err = ioutil.WriteFile(c.Race.FileToRead, []byte(text), 0644)
	if err != nil {
		log.Fatal("Failed to write file", err)
	}
	fmt.Println("File successfully updated with capitalized names!")
}