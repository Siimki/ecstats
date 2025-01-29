package db_test

import (
	"database/sql"
	"testing"
	_ "github.com/lib/pq"
	"ecstats/backend/db"
)

// Mock DB connection
func TestQueryRiderId(t *testing.T) {
	// Use the test database instead of production
	dbConn, err := sql.Open("postgres", "user=postgres password=admin port=5433 dbname=ecstats_test sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	defer dbConn.Close()

	// 🔥 Wipe test database before each test run (safe in test DB)
	_, err = dbConn.Exec(`TRUNCATE TABLE riders RESTART IDENTITY CASCADE;`)
	if err != nil {
		t.Fatalf("Failed to clear test database: %v", err)
	}

	// Insert mock data
	_, err = dbConn.Exec(`INSERT INTO riders (first_name, last_name, birth_year) VALUES ('Siim', 'TOMMY', 1997)`)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Call the function
	riderID := db.QueryRiderId(dbConn, "Siim", "TOMMY", 1997)

	// Expected ID should not be zero
	if riderID == 0 {
		t.Errorf("Expected a valid rider ID, got 0")
	}

	// Check non-existing rider
	riderID = db.QueryRiderId(dbConn, "John", "Doe", 2000)
	if riderID != 0 {
		t.Errorf("Expected 0 for a non-existing rider, got %d", riderID)
	}
}
