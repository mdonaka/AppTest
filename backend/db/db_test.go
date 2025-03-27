package db_test

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	dbpkg "backend/db"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T, data []dbpkg.Spices) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.Nil(t, err)

	_, err = db.Exec("CREATE TABLE spices (id INTEGER PRIMARY KEY, name TEXT, flavor TEXT, family TEXT)")
	assert.Nil(t, err)

	for _, entry := range data {
		_, err = db.Exec("INSERT INTO spices (id, name, flavor, family) VALUES (?, ?, ?, ?)", entry.Id, entry.Name, entry.Flavor, entry.Family)
		assert.Nil(t, err)
	}

	return db
}

// SelectAllが成功するかのテスト
func TestSelectUser_Success(t *testing.T) {
	testcases := []struct {
		name     string
		setup    []dbpkg.Spices
		expected []dbpkg.Spices
	}{
		{
			name: "select user",
			setup: []dbpkg.Spices{
				{0, "クミン", "スパイシー", "セリ科"},
			},
			expected: []dbpkg.Spices{{0, "クミン", "スパイシー", "セリ科"}},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db := setupTestDB(t, tt.setup)

			users, err := dbpkg.SelectAll(db)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, users)
		})
	}
}
