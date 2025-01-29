package main

import (
	"testing"
)

//go test -cover .
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out
//go test -v .    	to get a verbose output details of test
// go test -v -run TestMockDatabase
//go test -v -cover -run ^TestMockDatabase$
//go test -run Test_isPrime    to run individual test
//go test -run Test_alpha this will run all the test that are having alpha at their name


func TestMockDatabase(t *testing.T) {
	mockDB := &MockDatabase{}
	err := mockDB.Connect("") // Connection is a no-op in the mock
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}

	// Test Create
	err = mockDB.Create("users", map[string]interface{}{
		"id":   "1",
		"name": "John Doe",
		"age":  30,
	})
	if err != nil {
		t.Errorf("failed to create record: %v", err)
	}

	// Test Read
	user, err := mockDB.Read("users", "1")
	if err != nil {
		t.Errorf("failed to read record: %v", err)
	}
	if user["name"] != "John Doe" {
		t.Errorf("expected name 'John Doe', got %v", user["name"])
	}

	// Test Update
	err = mockDB.Update("users", "1", map[string]interface{}{
		"id":   "1",
		"name": "Jane Doe",
		"age":  25,
	})
	if err != nil {
		t.Errorf("failed to update record: %v", err)
	}

	updatedUser, err := mockDB.Read("users", "1")
	if err != nil {
		t.Errorf("failed to read updated record: %v", err)
	}
	if updatedUser["name"] != "Jane Doe" {
		t.Errorf("expected name 'Jane Doe', got %v", updatedUser["name"])
	}

	// Test Delete
	err = mockDB.Delete("users", "1")
	if err != nil {
		t.Errorf("failed to delete record: %v", err)
	}

	_, err = mockDB.Read("users", "1")
	if err == nil {
		t.Errorf("expected error when reading deleted record, got nil")
	}
}
