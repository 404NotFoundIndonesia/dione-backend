package domain

import (
	"context"
	"dione-backend/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
}
