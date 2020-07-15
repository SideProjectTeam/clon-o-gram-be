package users

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/SideProjectTeam/clon-o-gram-be/api/mngdb"
	"github.com/SideProjectTeam/clon-o-gram-be/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//TODO add loggers for errors
//TODO add better validations


//Register for user
func Register(c *gin.Context) {

	var body struct {
		Username string `validate:"required,min=3,max=15"`
		Password string `validate:"required,min=8"`
		Email    string `validate:"required,email"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.New().Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err!=nil{
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"username": body.Username})
}

//Login controller
func Login(c *gin.Context){
	type LoginCommand struct {
		Username string `json:"username" validate:"required,min=3,max=15"`
		Password string `json:"password" validate:"required,min=8"`
	}
	var body LoginCommand
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.New().Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := mngdb.Clonodb.Collection("Users").FindOne(mngdb.Ctx, bson.D{{"username", body.Username}}).Decode(&user); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password)); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"username or password is incorrect"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"token":"Login success: Need to return token here or in header"})
}

//Update is for PUT method 
func Update(c * gin.Context){

	//comment here for test 
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
	return

	var body struct {
		Fullname *string 			`json:"fullname"`
		Password *string 			`json:"password" validate:"min=3,max=15"`
		ID primitive.ObjectID       `json:"_id" validate:"required"`
		ProfilePicURL *string       `json:"profile_pic_url"`
		PrivacyLevel *bool          `json:"privacy_level"`
	}
	
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(body)
	
	//validate 
	if err := validator.New().Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Wait for token session
	//TODO add update on db
}