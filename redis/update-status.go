package redis

import (
	"fmt"
	"log"
)

func UpdateStatus(theater_id int64, time_slot_id int64, seat_id int, status string, shouldLog ...bool) {
	err := CLIENT.Set(
		fmt.Sprintf(
			"%d-%d-%d",
			theater_id,
			time_slot_id,
			seat_id),
		status,
		0).Err()

	if err != nil {
		log.Fatalf("Failed to set key: %v", err.Error())
	} else if len(shouldLog) == 0 || shouldLog[0] {
		log.Printf("Successfully set key: %d-%d-%d", theater_id, time_slot_id, seat_id)
	}
}
