package main

import (
	"log"
	"net/http"

	"github.com/syahjamal/go-medical/config"
	"github.com/syahjamal/go-medical/routes"
)

func main() {
	//Database
	config.InitDB()
	defer config.DB.Close()

	http.Handle("/", routes.Handlers())

	log.Println("Server Running localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
