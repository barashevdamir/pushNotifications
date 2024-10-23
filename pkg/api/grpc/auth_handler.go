package grpc

import (
	"context"
	"pushNotifications/pkg/service"
	authpb "pushNotifications/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	return &authpb.LoginResponse{Token: token}, nil
}

func (h *AuthHandler) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	valid, userID, err := h.authService.ValidateToken(req.Token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	return &authpb.ValidateTokenResponse{Valid: valid, UserId: userID}, nil
}
