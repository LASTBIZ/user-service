package user

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"lastbiz/user-service/pkg/user"
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
	Organization Organization `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Messengers   []Messenger  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u User) toGRPC() *user.User {
	var messengers []*user.Messenger
	for _, messenger := range u.Messengers {
		messengers = append(messengers, messenger.toGRPC())
	}

	return &user.User{
		Id:           u.ID,
		CreatedAt:    timestamppb.New(u.CreatedAt),
		UpdatedAt:    timestamppb.New(u.UpdatedAt),
		IsVerify:     u.IsVerify,
		Blocked:      u.Blocked,
		Email:        u.Email,
		Phone:        u.Phone,
		FirstName:    u.Firstname,
		LastName:     u.Lastname,
		FullName:     u.FullName,
		Role:         u.Role,
		Organization: u.Organization.toGRPC(),
		Messengers:   messengers,
	}
}

type Messenger struct {
	ID     uint32 `gorm:"primarykey"`
	Name   string
	Value  string
	UserId uint32
}

func (m Messenger) toGRPC() *user.Messenger {
	return &user.Messenger{
		Name:  m.Name,
		Value: m.Value,
	}
}

type Organization struct {
	ID     uint32 `gorm:"primarykey"`
	Name   string
	UserId uint32
}

func (o Organization) toGRPC() *user.Organization {
	return &user.Organization{
		Name: o.Name,
		Id:   o.UserId,
	}
}
