package modelUser

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         *string            `json:"Name" validate:"required,min=2,max=100"`
	Password     *string            `json:"password" validate:"required,min=6,max=100"`
	Email        *string            `json:"email" validate:"email,required"`
	PhoneNumber  *string            `json:"phoneNumber,omitempty"`
	Avator       *string            `json:"avator,omitempty"`
	AccessToken  *string            `json:"accessToken"`
	RefreshToken *string            `json:"refreshToken"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
	UserId       string             `json:"userId"`
}
