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
		flavor TEXT,
		family TEXT
	);
	`
	if _, err := conn.Exec(createTable); err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	insertStmt := `INSERT INTO spices (id, name, flavor, family) VALUES (?, ?, ?, ?);`
	for _, spice := range initialData {
		if _, err := conn.Exec(insertStmt, spice.Id, spice.Name, spice.Flavor, spice.Family); err != nil {
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
				{Id: 1, Name: "クミン", Flavor: "辛味", Family: "セリ科"},
				{Id: 2, Name: "コリアンダー", Flavor: "柑橘系", Family: "セリ科"},
			},
			expected: []db.Spices{
				{Id: 1, Name: "クミン", Flavor: "辛味", Family: "セリ科"},
				{Id: 2, Name: "コリアンダー", Flavor: "柑橘系", Family: "セリ科"},
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
