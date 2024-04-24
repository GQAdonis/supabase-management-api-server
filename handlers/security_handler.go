package handlers

import (
	"context"
	"tribemedia.io/m/services"
	api "tribemedia.io/m/supabase"
)

type securityHandler struct {
	api.SecurityHandler
	authService services.AuthService
}

func NewSecurityHandler(authService services.AuthService) api.SecurityHandler {
	return &securityHandler{
		authService: authService,
	}
}

func (h *securityHandler) HandleBearer(ctx context.Context, operationName string, t api.Bearer) (context.Context, error) {
	return h.SecurityHandler.HandleBearer(ctx, operationName, t)
}
