package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/syahjamal/go-medical/config"
	"github.com/syahjamal/go-medical/models"
	"golang.org/x/crypto/bcrypt"
)

//ErrorResponse struct
type ErrorResponse struct {
	Err string
}

//Error struct
type Error interface {
	Error() string
}

var db = config.DB

//CreateUser function
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)

	createdUser := db.Create(user)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdUser)

}
