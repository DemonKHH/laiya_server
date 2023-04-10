package response

import "time"

type LoginResponse struct {
	Name         *string   `json:"name" validate:"required,min=2,max=100"`
	Email        *string   `json:"email" validate:"email,required"`
	Avator       *string   `json:"avator"`
	Mac          *string   `json:"mac,omitempty"`
	AccessToken  *string   `json:"accessToken"`
	RefreshToken *string   `json:"refreshToken"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UserId       string    `json:"userId"`
}
