package auth

import (
	"net/http"
	"testing"
)

// Go looks for functions starting with "Test"
func TestGetAPIKey(t *testing.T) {
	
	// 1. Create fake data (Simulating a bad request)
	headers := http.Header{}
	headers.Set("Authorization", "Bearer super-secret-key-123")

	// 2. Run the function we are testing
	gotKey, err := GetAPIKey(headers)

	// 3. Check the results
	if err != nil {
		t.Fatalf("We expected no error, but got one: %v", err)
	}

	if gotKey != "super-secret-key-123" {
		t.Fatalf("We expected the key 'super-secret-key-123', but got: %s", gotKey)
	}
}
