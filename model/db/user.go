package db

import (
	"time"
)

type User struct {
	ID		 int32 `gorm:"primary_key"`
	Email    string `gorm:"type:varchar(50);UNIQUE;not null;" validate:"required"`
	Password string `gorm:"type:varchar(200);not null;" validate:"required"`
	Name     string `gorm:"type:varchar(50);UNIQUE;not null;" validate:"required"`
	Phone    string `gorm:"type:varchar(20);not null;" validate:"required"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;" validate:"required"`
	UpdatedAt time.Time	`gorm:"type:timestamp;"`
	LastLogin time.Time	`gorm:"type:timestamp;"`
}
