package routes

import (
	"encoding/json"
	"net/http"

	"github.com/syahjamal/go-medical/config"
	"github.com/syahjamal/go-medical/models"
)

//GetAllPasien function
func GetAllPasien(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// ctx, cancel := context.WithCancel(context.Background())

		// defer cancel()

	}
}

//CreatePasien function
func CreatePasien(w http.ResponseWriter, r *http.Request) {
	var pasien models.Pasien

	json.NewDecoder(r.Body).Decode(&pasien)

	config.DB.Create(&pasien)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pasien)
}
