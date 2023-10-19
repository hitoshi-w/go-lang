package model

type SignupRequest struct {
	Email string
	Password string
}

type SignupResponse struct {
	AccessToken string
	RefreshToken string
}