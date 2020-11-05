package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Pasien struct
type Pasien struct {
	gorm.Model
	NoRekamMedik string    `json:"no_rekam_medik"`
	NamaPasien   string    `json:"nama_pasien"`
	TglLahir     time.Time `json:"tgl_lahir"`
	JkPasien     string    `json:"jk_pasien"`
	AlamatPasien string    `json:"alamat_pasien" sql:"type:text;"`
	RekamMedik   string    `json:"rekam_medik" sql:"type:text;"`
	IDUser       string    `json:"id_user"`
}
