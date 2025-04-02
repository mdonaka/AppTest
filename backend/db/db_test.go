package db_test

import (
	"backend/db"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T, initialData []db.Spices) db.DB {
	t.Helper()
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open in-memory database: %v", err)
	}

	createTable := `
	CREATE TABLE spices (
		id INTEGER PRIMARY KEY,
		name TEXT,
    alias TEXT,
    taste TEXT,
		flavor TEXT,
		family TEXT,
    origin TEXT
	);
	`
	if _, err := conn.Exec(createTable); err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	insertStmt := `INSERT INTO spices (id, name, alias, taste, flavor, family, origin) VALUES (?, ?, ?, ?, ?, ?, ?)`
	for _, spice := range initialData {
		if _, err := conn.Exec(insertStmt, spice.Id, spice.Name, spice.Alias, spice.Taste, spice.Flavor, spice.Family, spice.Origin); err != nil {
			t.Fatalf("Failed to insert initial data: %v", err)
		}
	}

	return &db.SqliteDB{Conn: conn}
}

func TestSelectAll_Success(t *testing.T) {
	tests := []struct {
		name        string
		initialData []db.Spices
		expected    []db.Spices
	}{
		{
			name: "No error, multiple spices",
			initialData: []db.Spices{
				{Id: 1, Name: "クミン", Alias: "クミン", Taste: "辛味", Flavor: "辛味", Family: "セリ科", Origin: "インド"},
				{Id: 2, Name: "コリアンダー", Alias: "コリアンダー", Taste: "甘味", Flavor: "柑橘系", Family: "セリ科", Origin: "インド"},
			},
			expected: []db.Spices{
				{Id: 1, Name: "クミン", Alias: "クミン", Taste: "辛味", Flavor: "辛味", Family: "セリ科", Origin: "インド"},
				{Id: 2, Name: "コリアンダー", Alias: "コリアンダー", Taste: "甘味", Flavor: "柑橘系", Family: "セリ科", Origin: "インド"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			testDB := setupTestDB(t, tt.initialData)
			defer testDB.Close()

			actual, err := testDB.SelectAll()
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSelectByID_Success(t *testing.T) {
	tests := []struct {
		name        string
		initialData []db.Spices
		id          int
		expected    *db.Spices
	}{
		{
			name: "No error, single spice",
			initialData: []db.Spices{
				{Id: 1, Name: "クミン", Alias: "クミン", Taste: "辛味", Flavor: "辛味", Family: "セリ科", Origin: "インド"},
				{Id: 2, Name: "コリアンダー", Alias: "コリアンダー", Taste: "甘味", Flavor: "柑橘系", Family: "セリ科", Origin: "インド"},
			},
			id:       1,
			expected: &db.Spices{Id: 1, Name: "クミン", Alias: "クミン", Taste: "辛味", Flavor: "辛味", Family: "セリ科", Origin: "インド"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			testDB := setupTestDB(t, tt.initialData)
			defer testDB.Close()

			actual, err := testDB.SelectByID(tt.id)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSelectByID_Fail(t *testing.T) {
	tests := []struct {
		name          string
		initialData   []db.Spices
		id            int
		expectedError error
	}{
		{
			name:          "Error, spice not found",
			initialData:   []db.Spices{},
			id:            1,
			expectedError: sql.ErrNoRows,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			testDB := setupTestDB(t, tt.initialData)
			defer testDB.Close()

			_, err := testDB.SelectByID(tt.id)
			assert.EqualError(t, err, tt.expectedError.Error(), "Expected error: %v, got: %v", tt.expectedError, err)
		})
	}
}
