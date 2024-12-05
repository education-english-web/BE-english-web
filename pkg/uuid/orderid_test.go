package uuid

import (
	"strconv"
	"testing"
)

func TestOrderID(t *testing.T) {
	orderID := OrderID()

	// Check if orderID is not zero
	if orderID == 0 {
		t.Error("OrderID is zero")
	}

	// Additional check: Ensure orderID is within expected range (e.g., non-negative)
	if orderID < 0 {
		t.Errorf("OrderID is negative: %d", orderID)
	}

	// Optional: Check the length of the orderID (since it's a subset of Unix timestamp in nanoseconds)
	orderIDStr := strconv.FormatInt(orderID, 10)
	expectedLength := 9
	if len(orderIDStr) != expectedLength {
		t.Errorf("OrderID has incorrect length: got %d, want %d", len(orderIDStr), expectedLength)
	}
}
