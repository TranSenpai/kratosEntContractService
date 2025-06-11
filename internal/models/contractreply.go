package requestModel

import "time"

type ReplyContract struct {
	Id                   *uint64
	StudentCode          *string
	FirstName            *string
	LastName             *string
	MiddleName           *string
	Email                *string
	Phone                *string
	Gender               *uint32
	Dob                  *time.Time
	Address              *string
	Avatar               *string
	RoomId               *string
	IsActive             *bool
	Sign                 *string
	RegistryAt           *time.Time
	NotificationChannels *uint32
}
