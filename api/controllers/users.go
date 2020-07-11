package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	
	"gopkg.in/go-playground/validator.v9"
	"github.com/SideProjectTeam/clon-o-gram-be/api/mngdb"
	"github.com/SideProjectTeam/clon-o-gram-be/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO add better validations
type CreateUserBody struct {
	Username string `validate:"required,min=3,max=15"`
	Password string `validate:"required,min=8"`
	Email    string `validate:"required,email"`
}



func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var body CreateUserBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if err := validator.New().Struct(body); err != nil {
		http.Error(w, err.Error(), 400)
	}

	// need password hashing
	user := &models.User{
		ID: primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username: body.Username,
		Password: body.Password, 
		Email: body.Email,
	}
	//Create
	res,err := mngdb.Clonodb.Collection("Users").InsertOne(mngdb.Ctx, user)
	fmt.Println(res)
	if err!=nil{
		fmt.Println(err)
	}

}
