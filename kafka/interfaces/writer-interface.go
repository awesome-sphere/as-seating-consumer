package interfaces

type WriterInterface struct {
	TheaterID  int64  `json:"theater_id"`
	TimeSlotID int64  `json:"time_slot_id"`
	SeatID     int    `json:"seat_number"`
	SeatStatus string `json:"seat_status"`
}
