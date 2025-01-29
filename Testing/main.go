package main

import (
	s "db/testing/simple"
	"errors"
	"fmt"
)

func main() {
	a := []int{4, 7, 8, 9, 10}
	fmt.Println(s.FetchMaxNumber(a))
}

// MockDatabase simulates a database in memory for testing purposes.
type MockDatabase struct {
	data map[string]map[string]interface{} // Simulates a table with rows
}

func (m *MockDatabase) Connect(connectionString string) error {
	// No actual connection needed for a mock
	m.data = make(map[string]map[string]interface{}) // Initialize mock data
	return nil
}

func (m *MockDatabase) Create(table string, data map[string]interface{}) error {
	if _, exists := m.data[table]; !exists {
		m.data[table] = make(map[string]interface{}) // Initialize table
	}

	// Assume the "id" field is the key
	id, ok := data["id"].(string)
	if !ok || id == "" {
		return errors.New("missing or invalid 'id' field")
	}

	m.data[table][id] = data
	return nil
}

func (m *MockDatabase) Read(table string, id string) (map[string]interface{}, error) {
	if rows, exists := m.data[table]; exists {
		if row, exists := rows[id]; exists {
			return row.(map[string]interface{}), nil
		}
	}
	return nil, errors.New("record not found")
}

func (m *MockDatabase) Update(table string, id string, data map[string]interface{}) error {
	if rows, exists := m.data[table]; exists {
		if _, exists := rows[id]; exists {
			m.data[table][id] = data
			return nil
		}
	}
	return errors.New("record not found")
}

func (m *MockDatabase) Delete(table string, id string) error {
	if rows, exists := m.data[table]; exists {
		if _, exists := rows[id]; exists {
			delete(rows, id)
			return nil
		}
	}
	return errors.New("record not found")
}
