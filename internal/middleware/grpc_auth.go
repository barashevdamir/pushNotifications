package middleware

import (
	"context"
	"pushNotifications/pkg/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryAuthInterceptor(authService service.AuthService) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Пропускаем методы авторизации
		if info.FullMethod == "/auth_service.AuthService/Login" {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
		}

		var token string
		if val, ok := md["authorization"]; ok && len(val) > 0 {
			token = val[0]
		} else {
			return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
		}

		valid, userID, err := authService.ValidateToken(token)
		if err != nil || !valid {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		// Добавляем userID в контекст
		ctx = context.WithValue(ctx, "userID", userID)

		return handler(ctx, req)
	}
}
