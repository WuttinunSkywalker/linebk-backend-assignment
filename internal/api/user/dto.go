package user

import "time"

type UserResponse struct {
	UserID    string    `json:"user_id" example:"25799236-5dbc-411a-8641-b8336e088b9b"`
	Name      string    `json:"name" example:"Misty Von-Lemke"`
	Image     *string   `json:"image" example:"https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/47.jpg"`
	CreatedAt time.Time `json:"created_at" example:"2021-03-07T04:03:06Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-02-08T02:06:17Z"`
}

func NewUserResponse(user *User) *UserResponse {
	return &UserResponse{
		UserID:    user.UserID,
		Name:      user.Name,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type UserGreetingResponse struct {
	UserID    string    `json:"user_id" example:"25799236-5dbc-411a-8641-b8336e088b9b"`
	Greeting  string    `json:"greeting" example:"Hello Misty Von-Lemke"`
	CreatedAt time.Time `json:"created_at" example:"2021-03-07T10:31:51Z"`
}

func NewUserGreetingResponse(greeting *UserGreeting) *UserGreetingResponse {
	return &UserGreetingResponse{
		UserID:    greeting.UserID,
		Greeting:  greeting.Greeting,
		CreatedAt: greeting.CreatedAt,
	}
}

type UserPreviewResponse struct {
	Name  string  `json:"name" example:"Misty Von-Lemke"`
	Image *string `json:"image" example:"https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/47.jpg"`
}

func NewUserPreviewResponse(user *User) *UserPreviewResponse {
	return &UserPreviewResponse{
		Name:  user.Name,
		Image: user.Image,
	}
}
