package models

import (
	"time"
)

//User struct
type User struct {
	Pasien    []Pasien
	IDUser    uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
