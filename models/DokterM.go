package models

import "github.com/jinzhu/gorm"

//Dokter struct
type Dokter struct {
	gorm.Model
	IDDokter     string
	NamaDokter   string
	JkDokter     string
	AlamatDokter string `sql:"type:text;"`
	Spesialis    string
}
