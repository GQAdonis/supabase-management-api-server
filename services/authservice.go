package services

import "github.com/supabase-community/gotrue-go"

type AuthService interface{}

type authService struct {
	goTrueClient gotrue.Client
}

func NewAuthService(goTrueClient gotrue.Client) AuthService {
	return &authService{
		goTrueClient: goTrueClient,
	}
}
