package db

import "time"

type Transaction struct {
	Serial string `gorm:"type:varchar(150);UNIQUE;not null;" validate:"required"`
	UserId int32 `gorm:"type:integer;not null;" validate:"required"`
	OperType int `gorm:"type:smallint;not null;" validate:"required"`
	Amount int64 `gorm:"type:integer;not null;" validate:"required"`
	Balance int64 `gorm:"type:integer;"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;" validate:"required"`
}
