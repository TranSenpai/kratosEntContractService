package requestModel

type TotalContractsEachRoom struct {
	Total  uint8  `json:"total"`
	RoomID string `json:"room_id"`
}
