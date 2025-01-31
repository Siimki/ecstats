package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const RaceId = 57
const FileToRead = "../results/BOSCH/output.txt"
const Year = 2024

var PopularNamesToCapitalize = []string{
	"Puri", "Ruder", "Pärn", "Rannaveer", "Jakobson", "Rebane", "Tuula-Fjodorov", "Valge-Rebane", "Evendi",
	"Grünberg", "Kiigemägi", "Mõttus", "Vendelin", "Varik", "Truus", "Meidla", "Ermane-Marcenko",
	"Juutinen", "Saar", "Põldmaa", "Vernik", "Rei", "Käärmann-Liive", "Algpeus",
	"Glinka-Valkrusman", "Paap", "Luuka", "Nigul", "Jürjenberg", "Paldra", "Mollberg", "Mauer",
	"Orav", "REInu", "Reinu", "Palu", "Tilk", "Rooba", "Mosov-Hallik", "Allik", "Mäemurd",
	"Taal", "Baumann", "Metsmaa", "Schildhauer", "Puhke",
}

// Config structure
type Config struct {
	Race struct {
		ID int `yaml:"id"`
	} `yaml:"race"`

	File struct {
		Path string `yaml:"path"`
	} `yaml:"file"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

// LoadConfig reads the config.yaml file
func LoadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	return &config, nil
}
