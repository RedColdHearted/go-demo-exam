package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseReservation struct {
	ClientName string `json:"client_name" gorm:"type:varchar(120);not null" fake:"{firstname} {lastname}"`
	ClientPhone string `json:"client_phone" gorm:"type:varchar(120);not null" fake:"{phone}"`
	Whishes *string `json:"whishes" gorm:"type:text" fake:"{sentence:6}"`
	Address string `json:"address" gorm:"type:varchar(120)" fake:"{address}"`
	RoomID uint `json:"room_id" gorm:"not null" fake:"{number:1,100}"`
	CheckIn time.Time `json:"check_in" gorm:"not null" fake:"{year}-{month}-{day}" format:"2006-1-2"`
	CheckOut time.Time `json:"check_out" gorm:"not null" fake:"{year}-{month}-{day}" format:"2006-1-2"`
}

type CreateReservation struct{
	BaseReservation
}

type UpdateReservation struct{
	BaseReservation
}

type Reservation struct{
	ID        uint `json:"-" gorm:"primarykey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	BaseReservation
}

