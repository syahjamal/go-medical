package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syahjamal/go-medical/controller"
)

//Handlers function
func Handlers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controller.TestAPI).Methods("GET")
	router.HandleFunc("/register", controller.CreateUser).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")

	router.HandleFunc("/pasien/create", CreatePasien).Methods("POST")

	return router
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
