package auth

type LoginWithPinRequest struct {
	UserID string `json:"user_id" binding:"required" example:"0befecd8-fccb-417e-aa0a-1a23c021f413"`
	Pin    string `json:"pin" binding:"required,len=6" example:"123456"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJsaW5lYmstYmFja2VuZC1hc3NpZ25tZW50Iiwic3ViIjoidGVzdC11c2VyIiwiZXhwIjoxNzUwNzU5MTIzLCJpYXQiOjE3NTA3MTU5MjN9.NSuzaSM6CdrnNtcsrybBsn_2UhGpOlR5g1fOWuHYrzM"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJsaW5lYmstYmFja2VuZC1hc3NpZ25tZW50Iiwic3ViIjoidGVzdC11c2VyIiwiZXhwIjoxNzUwODAyMzIzLCJpYXQiOjE3NTA3MTU5MjN9.hM7mwFmaXW0CZj1JYg44ic-94g3Ngbpg4P-7SGAcXtI"`
}
