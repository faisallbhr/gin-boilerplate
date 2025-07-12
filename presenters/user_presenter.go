package presenters

import "github.com/faisallbhr/gin-boilerplate/models"

type UserResponse struct {
	Id           uint    `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Token        *string `json:"token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
}

func FormatUsers(users []models.User) []UserResponse {
	result := make([]UserResponse, len(users))
	for i, user := range users {
		result[i] = UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}
	}

	return result
}
