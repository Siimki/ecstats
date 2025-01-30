package db_test

import (
	"testing"
	"ecstats/backend/dataclean"
)

// Test cases for extracting first and last names
func TestExtractNames(t *testing.T) {
	tests := []struct {
		fullName   string
		firstName  string
		lastName   string
	}{
		{"Martin Mihkel MEIDLA", "Martin Mihkel", "MEIDLA"},
		{"Siim KISKONEN TAMM", "Siim", "KISKONEN TAMM"},
		{"John DOE", "John", "DOE"},
		{"Peeter TARVIS", "Peeter", "TARVIS"},
		{"Andres VELTSON", "Andres", "VELTSON"},
		{"Mihkel Tamm", "Mihkel Tamm", ""}, // No uppercase last name
		{"Karl Jüri KISKONEN-TAMM", "Karl Jüri", "KISKONEN-TAMM"},
		{"Markus VÄLI", "Markus", "VÄLI"},
		{"Heiki LÕHMUS KASK", "Heiki", "LÕHMUS KASK"},
		{"Indrek LEPPIK", "Indrek", "LEPPIK"},
		{"Rai-Ner OTS", "Rai-Ner", "OTS"},
		{"Rainer OTS-US", "Rainer", "OTS-US"},
		{"Rainer OTS", "Rainer", "OTS"},
	}

	for _, test := range tests {
		firstName, lastName := dataclean.ExtractNames(test.fullName)
		if firstName != test.firstName || lastName != test.lastName {
			t.Errorf("extractNames(%q) = (%q, %q); want (%q, %q)",
				test.fullName, firstName, lastName, test.firstName, test.lastName)
		}
	}
	
}
