package test

import (
	"flipSensorServer/database"
	"testing"
)

func TestInsertEntry(t *testing.T) {
	db := database.SetupTestDB(t)
	defer db.Close()

	entry := database.DataEntry{
		TimeStamp:   "2025-01-01T10:00:00Z",
		Temperature: 25.5,
		GasLevel:    0.8,
	}

	_, err := database.InsertEntry(db, entry)
	if err != nil {
		t.Fatalf("insertEntry failed: %v", err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM entries").Scan(&count)
	if err != nil {
		t.Fatalf("query failed: %v", err)
	}

	if count != 1 {
		t.Fatalf("expected 1 entry, got %d", count)
	}
}

func TestGetEntries(t *testing.T) {
	db := database.SetupTestDB(t)
	defer db.Close()

	entriesToInsert := []database.DataEntry{
		{"2025-01-01T10:00:00Z", 20.0, 0.5},
		{"2025-01-01T11:00:00Z", 21.0, 0.6},
	}

	for _, e := range entriesToInsert {
		_, err := database.InsertEntry(db, e)
		if err != nil {
			t.Fatalf("insert failed: %v", err)
		}
	}

	entries, err := database.GetEntries(db)
	if err != nil {
		t.Fatalf("getEntries failed: %v", err)
	}

	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
}
