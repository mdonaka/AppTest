package api_test

import (
	"backend/api"
	"backend/db"
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockDB struct {
	data []db.Spices
	err  error
}

func (m *MockDB) SelectAll() ([]db.Spices, error) {
	return m.data, m.err
}

func (m *MockDB) SelectByID(id int) (*db.Spices, error) {
	if m.err != nil {
		return nil, m.err
	}
	for _, spice := range m.data {
		if spice.Id == id {
			return &spice, nil
		}
	}
	return nil, nil
}

func (m *MockDB) Close() {}

func TestDataHandler_Success(t *testing.T) {
	tests := []struct {
		name           string
		mockDB         *MockDB
		expectedStatus int
		expectedData   []db.Spices
	}{
		{
			name: "Success",
			mockDB: &MockDB{
				data: []db.Spices{
					{Id: 1, Name: "クミン", Flavor: "辛味", Family: "セリ科"},
					{Id: 2, Name: "コリアンダー", Flavor: "柑橘系", Family: "セリ科"},
				},
				err: nil,
			},
			expectedStatus: http.StatusOK,
			expectedData: []db.Spices{
				{Id: 1, Name: "クミン", Flavor: "辛味", Family: "セリ科"},
				{Id: 2, Name: "コリアンダー", Flavor: "柑橘系", Family: "セリ科"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			handler := &api.HandlerWithDB{DB: tt.mockDB}
			req, err := http.NewRequest("GET", "/data", nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.DataHandler(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			var actual []db.Spices
			err = json.NewDecoder(rr.Body).Decode(&actual)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedData, actual)
		})
	}
}

func TestDataHandler_Failed(t *testing.T) {
	tests := []struct {
		name           string
		mockDB         *MockDB
		expectedStatus int
	}{
		{
			name: "DB Error",
			mockDB: &MockDB{
				data: nil,
				err:  assert.AnError,
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			handler := &api.HandlerWithDB{DB: tt.mockDB}
			req, err := http.NewRequest("GET", "/data", nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.DataHandler(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}

func TestCheckSpiceHandler_Success(t *testing.T) {
	tests := []struct {
		name           string
		mockDB         *MockDB
		queryParams    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "true1",
			mockDB: &MockDB{
				data: []db.Spices{
					{Id: 1, Name: "クミン", Flavor: "辛味", Family: "セリ科"},
				},
				err: nil,
			},
			queryParams:    "id=1&name=クミン&flavor=辛味&family=セリ科",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"match":true}`,
		},
		{
			name: "true2",
			mockDB: &MockDB{
				data: []db.Spices{
					{Id: 1, Name: "クミン1", Flavor: "辛味", Family: "セリ科"},
					{Id: 2, Name: "クミン2", Flavor: "辛味", Family: "セリ科"},
				},
				err: nil,
			},
			queryParams:    "id=2&name=クミン2&flavor=辛味&family=セリ科",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"match":true}`,
		},
		{
			name: "false1",
			mockDB: &MockDB{
				data: []db.Spices{
					{Id: 1, Name: "クミン", Flavor: "辛味", Family: "セリ科"},
				},
				err: nil,
			},
			queryParams:    "id=1&name=コリアンダー&flavor=辛味&family=セリ科",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"match":false}`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			handler := &api.HandlerWithDB{DB: tt.mockDB}
			req, err := http.NewRequest("GET", "/check?"+tt.queryParams, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.CheckSpiceHandler(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestCheckSpiceHandler_Failed(t *testing.T) {
	tests := []struct {
		name           string
		mockDB         *MockDB
		queryParams    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Missing one or more parameters",
			mockDB: &MockDB{
				data: nil,
				err:  nil,
			},
			queryParams:    "id=1",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Missing one or more parameters\n",
		},
		{
			name: "Invalid id parameter",
			mockDB: &MockDB{
				data: nil,
				err:  nil,
			},
			queryParams:    "id=abc&name=クミン&flavor=辛味&family=セリ科",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid id parameter\n",
		},
		{
			name: "Spice not found",
			mockDB: &MockDB{
				data: nil,
				err:  sql.ErrNoRows,
			},
			queryParams:    "id=999&name=クミン&flavor=辛味&family=セリ科",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Spice not found\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			handler := &api.HandlerWithDB{DB: tt.mockDB}
			req, err := http.NewRequest("GET", "/check?"+tt.queryParams, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.CheckSpiceHandler(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}
