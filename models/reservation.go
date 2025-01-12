package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseReservation struct {
	ClientName string `json:"client_name" gorm:"type:varchar(120);not null"`
	ClientPhone string `json:"client_phone" gorm:"type:varchar(120);not null"`
	Whishes *string `json:"whishes" gorm:"type:text"`
	Address string `json:"address" gorm:"type:varchar(120)"`
	RoomID uint `json:"room_id" gorm:"not null"`
	CheckIn time.Time `json:"check_in" gorm:"not null"`
	CheckOut time.Time `json:"check_out" gorm:"not null"`
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

