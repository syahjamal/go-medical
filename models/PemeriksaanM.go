package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Pemeriksaan struct
type Pemeriksaan struct {
	gorm.Model
	IDPemeriksaan  string
	IDPasien       string
	Keluhan        string
	TglPemeriksaan time.Time
	Status         string
	Diagnosa       string `sql:"type:text;"`
	ResepObat      string `sql:"type:text;"`
}
