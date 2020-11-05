package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Booking struct
type Booking struct {
	gorm.Model
	KodeBookin string
	TglBooking time.Time
	NoAntrian  string
	NamaPasien string
	NamaDokter string
	KodeJadwal string
	Status     string
}
