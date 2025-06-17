package entities

import (
	"time"
)

// type User interface {
// 	RegistryRoom()
// }

// varchar in MySQL use 1 byte (2^8 - 1 = 255) to store length of record and another one - four byte each character
// the reason why subtract 1 for store the & to connect the string
// so if the record's length is < 255, we should declare the length is 255 to take the most of capacity of bit length
// and if the record's length is > 255 and < 510 use the size 510 -> length = 2^n - 1
type Contract struct {
	ID                   uint64     `gorm:"type:bigint UNSIGNED; primaryKey"`
	StudentCode          string     `gorm:"type:char(10); index:idx_student_code"`
	FirstName            string     `gorm:"type:varchar(255); index:idx_student_info, priority:2"`
	LastName             string     `gorm:"type:varchar(255); index:idx_student_info, priority:1"`
	MiddleName           string     `gorm:"type:varchar(255);"`
	Email                string     `gorm:"type:varchar(128); index:idx_email"`
	Sign                 string     `gorm:"type:varchar(255); index:idx_sign"` // Because bcrypt algorithm return 60 characteres
	Phone                string     `gorm:"type:char(10); index: idx_phone"`
	Gender               uint8      `gorm:"type:tinyint UNSIGNED;"`
	DOB                  *time.Time `gorm:"type:TIMESTAMP;"`
	Address              string     `gorm:"type:varchar(255)"`
	Avatar               []byte     `gorm:"type:mediumblob"` // Should store varchar(1024) (link img)
	IsActive             bool       `gorm:"type:boolean; index: idx_is_active"`
	RegistryAt           time.Time  `gorm:"type:timestamp;"`
	RoomID               string     `gorm:"type:char(5); index: idx_room_id"`
	NotificationChannels uint8      `gorm:"type:tinyint UNSIGNED;"`
}
