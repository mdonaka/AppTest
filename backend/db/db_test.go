package db_test

import (
	"database/sql"
	"testing"

	dbpkg "backend/db"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(data []dbpkg.User) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
	if err != nil {
		return nil, err
	}

	for _, entry := range data {
		_, err = db.Exec("INSERT INTO users (id, name, age) VALUES (?, ?, ?)", entry.Id, entry.Name, entry.Age)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func TestSelectUser_Success(t *testing.T) {
	testcases := []struct {
		name     string
		setup    []dbpkg.User
		expected []dbpkg.User
	}{
		{
			name: "select user",
			setup: []dbpkg.User{
				{1, "Alice", 20},
			},
			expected: []dbpkg.User{{1, "Alice", 20}},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, err := setupTestDB(tt.setup)
			if err != nil {
				t.Fatalf("Failed setup test db: %v", err)
			}
			defer db.Close()

			users, err := dbpkg.SelectAll(db)
			if err != nil {
				t.Fatalf("Failed to select user: %v", err)
			}
			if len(users) != len(tt.expected) {
				t.Fatalf("expected %v, but got %v", tt.expected, users)
			}
			for i, user := range users {
				if user.Id != tt.expected[i].Id {
					t.Errorf("expected %v, but got %v", tt.expected[i].Id, user.Id)
				}
				if user.Name != tt.expected[i].Name {
					t.Errorf("expected %v, but got %v", tt.expected[i].Name, user.Name)
				}
				if user.Age != tt.expected[i].Age {
					t.Errorf("expected %v, but got %v", tt.expected[i].Age, user.Age)
				}
			}
		})
	}
}
func TestSelectUser_Failed(t *testing.T) {
	testcases := []struct {
		name          string
		setup         []dbpkg.User
		expectedError error
	}{
		{
			name: "select user",
			setup: []dbpkg.User{
				{1, "Alice", 20},
			},
			expectedError: nil,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, err := setupTestDB(tt.setup)
			if err != nil {
				t.Fatalf("Failed setup test db: %v", err)
			}
			defer db.Close()
			if err != tt.expectedError {
				t.Fatalf("expected %v, but got %v", tt.expectedError, err)
			}
		})
	}
}
