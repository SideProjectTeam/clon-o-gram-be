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


type CreateUserBody struct {
	Username string `validate:"required,min=3,max=15"`
	Password string `validate:"required,min=8"`
	Email    string `validate:"required,email"`
}

//Register
func Register(c *gin.Context) {

	var body CreateUserBody
	//c.BindJSON(&body)

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
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var body LoginCommand
	if err := c.ShouldBindJSON(&body); err != nil {
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