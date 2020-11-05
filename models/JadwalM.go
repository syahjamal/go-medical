package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Jadwal Struct
type Jadwal struct {
	gorm.Model
	KodeJadwal  string
	Hari        string
	JadwalMulai time.Time
	JadwalAkhir time.Time
}
