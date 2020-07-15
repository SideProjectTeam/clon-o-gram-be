
package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

//User model
type User struct {
    ID primitive.ObjectID        `bson:"_id"`
    CreatedAt time.Time          `bson:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at"`
    Username  string             `bson:"username"`
	Password string              `bson:"password"`
    Email string				 `bson:"email"`
    FullName  string             `bson:"username"`
    ProfilePicURL string         `bson:"profile_pic_url"`
    PrivacyLevel bool            `bson:"privacy_level"`
}