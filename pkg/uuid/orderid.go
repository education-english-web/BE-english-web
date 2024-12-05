package uuid

import (
	"fmt"
	"strconv"
	"time"
)

func OrderID() int64 {
	// Get the current time
	t := time.Now()

	// Convert the time to Unix timestamp
	timestamp := t.UnixNano() // Use UnixNano for more digits

	// Convert the Unix timestamp to a string
	order_id_str := fmt.Sprintf("%d", timestamp)

	// Take the first 16 digits (or however many you prefer)
	if len(order_id_str) > 9 {
		order_id_str = order_id_str[:9]
	}

	// Convert the string to int64
	order_id, err := strconv.ParseInt(order_id_str, 10, 64)
	if err != nil {
		return 0
	}
	return order_id
}
