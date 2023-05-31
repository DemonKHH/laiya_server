package response

import "time"

type LoginResponse struct {
	Name         *string   `json:"name" validate:"required,min=2,max=100"`
	Email        *string   `json:"email" validate:"email,required"`
	Avator       *string   `json:"avator"`
	AccessToken  *string   `json:"accessToken"`
	RefreshToken *string   `json:"refreshToken"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UserId       string    `json:"userId"`
	Permissions  []string  `json:"permissions"`
}

type LoginResponseWithUser struct {
	Name        *string  `json:"name" validate:"required,min=2,max=100"`
	Email       *string  `json:"email" validate:"email,required"`
	UserId      string   `json:"userId"`
	Permissions []string `json:"permissions"`
}
