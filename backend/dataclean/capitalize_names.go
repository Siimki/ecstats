package dataclean

import (
	"fmt"
	"io/ioutil"
	"strings"
	"ecstats/backend/config"
	"log"
)

func CapitalizePopularLastNames() {

	content, err := ioutil.ReadFile(config.FileToRead)
	if err != nil {
		log.Fatal("Cant read file in capitalize")
		return
	}

	text := string(content)
	for _, name := range config.PopularNamesToCapitalize {
		text = strings.ReplaceAll(text, name, strings.ToUpper(name))
	}

	err = ioutil.WriteFile(config.FileToRead, []byte(text), 0644)
	if err != nil {
		log.Fatal("Failed to write file", err)
	}
	fmt.Println("File successfully updated with capitalized names!")
}