package requestModel

import "time"

type ReplyContract struct {
	Id                   int
	StudentCode          string
	FirstName            string
	LastName             string
	MiddleName           string
	Email                string
	Phone                string
	Gender               uint32
	Dob                  time.Time
	Address              string
	Avatar               string
	RoomId               string
	IsActive             bool
	Sign                 string
	NotificationChannels uint32
	RegistryAt           time.Time
}
