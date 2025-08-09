package service

import (
	"context"
	"database/sql"
	"dione-backend/domain"
	"dione-backend/dto"
	"dione-backend/internal/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type authService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuthService(conf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return &authService{conf: conf, userRepository: userRepository}
}

func (a authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if user.ID == "" {
		return dto.LoginResponse{}, errors.New("user not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return dto.LoginResponse{}, errors.New("invalid password")
	}

	claim := jwt.MapClaims{
		"id": user.ID, "role": user.Role,
		"exp": time.Now().Add(time.Minute * time.Duration(a.conf.Jwt.Exp)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return dto.LoginResponse{}, errors.New("failed to sign token")
	}

	return dto.LoginResponse{Token: t}, nil
}

func (a authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	if user, _ := a.userRepository.FindByEmail(ctx, req.Email); user.ID != "" {
		return dto.RegisterResponse{}, errors.New("the email is already registered")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterResponse{}, errors.New("failed to hash password")
	}

	user := domain.User{
		ID:        uuid.NewString(),
		Name:      req.Name,
		Email:     req.Email,
		Phone:     "",
		Role:      domain.UserRoleUser,
		Password:  string(password),
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	if err = a.userRepository.Save(ctx, &user); err != nil {
		return dto.RegisterResponse{}, err
	}

	claim := jwt.MapClaims{
		"id": user.ID, "role": user.Role,
		"exp": time.Now().Add(time.Minute * time.Duration(a.conf.Jwt.Exp)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return dto.RegisterResponse{}, errors.New("failed to sign token")
	}

	return dto.RegisterResponse{
		Token: t,
	}, nil
}
