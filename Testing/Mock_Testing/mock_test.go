package mocktesting

import "testing"

// Mock implementation
type MockDatabase struct{}

func (m *MockDatabase) GetData(id string) string {
	return "mock data"
}

// Test with the mock
func TestFetchUserData(t *testing.T) {
	mockDB := &MockDatabase{}
	result := FetchUserData(mockDB, "123")
	if result != "mock data" {
		t.Errorf("Expected 'mock data', but got %s", result)
	}
}
