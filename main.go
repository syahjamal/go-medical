package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syahjamal/go-medical/config"
	"github.com/syahjamal/go-medical/controller"
	"github.com/syahjamal/go-medical/routes"
)

func main() {
	//Database
	config.InitDB()
	defer config.DB.Close()

	router := mux.NewRouter()

	router.HandleFunc("/register", controller.CreateUser).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")

	router.HandleFunc("/pasien/create", routes.CreatePasien).Methods("POST")

	log.Println("Server Running localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
