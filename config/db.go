package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql
	"github.com/syahjamal/go-medical/models"
)

//DB variable
var DB *gorm.DB

//InitDB function
func InitDB() {
	var err error

	// dsn := "user=postgres password=syahjamal dbname=go_medic port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	// DB, err := gorm.Open("postgres", dsn)
	DB, err = gorm.Open("mysql", "root:@/go_medic?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.User{}, &models.Pasien{}, &models.Dokter{}, &models.Jadwal{}, &models.Booking{}, &models.Pemeriksaan{})
}
