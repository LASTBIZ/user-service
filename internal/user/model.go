package user

import (
	"time"
)

type User struct {
	ID           uint32 `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsVerify     bool `gorm:"default:false"`
	Firstname    string
	Lastname     string
	FullName     string `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
	Blocked      bool   `gorm:"default:false"`
	Role         string
	Email        string
	Phone        string
	Organization Organization
	Messengers   []Messenger
}

type Messenger struct {
	Name   string
	Value  string
	UserId uint32
}

type Organization struct {
	ID     uint32 `gorm:"primarykey"`
	Name   string
	UserId uint32
}
