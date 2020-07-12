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
	"golang.org/x/crypto/bcrypt"
)

//TODO add loggers for errors
//TODO add better validations


type CreateUserBody struct {
	Username string `validate:"required,min=3,max=15"`
	Password string `validate:"required,min=8"`
	Email    string `validate:"required,email"`
}


//GetUsers Dumb method
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

//RegisterUser is controller function for register a user
func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var body CreateUserBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := validator.New().Struct(body); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.MinCost) 
	if err != nil {
		fmt.Println(err)
		return
    }
	user := &models.User{
		ID: primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username: body.Username,
		Password: string(hash), 
		Email: body.Email,
	}

	//Create
	_,err = mngdb.Clonodb.Collection("Users").InsertOne(mngdb.Ctx, user)
	if err!=nil{
		http.Error(w,err.Error(),400)
		return
	}

	if err!=nil{
		fmt.Println(err)
	}

	fmt.Fprintf(w,"User Created %s",body.Username)
}
