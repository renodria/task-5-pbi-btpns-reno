package models

import (
	"time"
)

type User struct {
	ID        int16     `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Email     string    `gorm:"type:text" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Photos    []Photos  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE" json:"photos"`
}
