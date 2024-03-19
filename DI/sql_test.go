package sql

import (
	"database/sql"
	"testing"
)

type mockDB struct {
	query        string
	lastInsertId int64
	rowsAffected int64
}

// LastInsertId implements sql.Result.
func (m *mockDB) LastInsertId() (int64, error) {
	return m.lastInsertId, nil
}

// RowsAffected implements sql.Result.
func (m *mockDB) RowsAffected() (int64, error) {
	return m.rowsAffected, nil
}

func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.query = query
	return m, nil
}

func TestExecQuery(t *testing.T) {
	mock := &mockDB{
		rowsAffected: 32,
	}

	r, _ := execQuery(mock, "SELECT * FROM sql")

	if mock.query != "SELECT * FROM sql" {
		t.Errorf("got %v want %v", mock.query, "SELECT * FROM sql")
	}

	if r != 32 {
		t.Errorf("got %v want %v", r, 32)
	}
}
