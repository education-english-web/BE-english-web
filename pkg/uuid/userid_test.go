package uuid

import (
	"testing"
)

func TestCreateUserID(t *testing.T) {
	email := "example@email.com"
	userID := CreatUserID(email)

	// Check if UserID is empty
	if userID == "" {
		t.Error("UserID is empty")
	}

	// Additional check: Ensure UserID is of the correct length
	expectedLength := 36 // UUID is typically 36 characters long
	if len(userID) != expectedLength {
		t.Errorf("UserID has incorrect length: got %d, want %d", len(userID), expectedLength)
	}
}
