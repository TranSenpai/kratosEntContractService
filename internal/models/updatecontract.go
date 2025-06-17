package requestModel

import (
	"time"
)

type UpdateContract struct {
	ID                   *uint64
	StudentCode          *string
	FirstName            *string
	LastName             *string
	MiddleName           *string
	Email                *string
	Sign                 *string
	Phone                *string
	Gender               *uint32
	Dob                  *time.Time
	Address              *string
	Avatar               *string
	IsActive             *bool
	RoomID               *string
	NotificationChannels *uint32
}
