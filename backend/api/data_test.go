package api_test

import (
	"backend/api"
	"backend/db"
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
			handler := &api.DataHandlerStruct{DB: tt.mockDB}
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
			handler := &api.DataHandlerStruct{DB: tt.mockDB}
			req, err := http.NewRequest("GET", "/data", nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.DataHandler(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}
