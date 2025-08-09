package service

import (
	"context"
	"dione-backend/domain"
	"dione-backend/dto"
	"net/url"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) Show(ctx context.Context, id string) (dto.UserData, error) {
	user, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		return dto.UserData{}, err
	}

	avatarUrl := user.AvatarPath
	if avatarUrl == "" {
		avatarUrl = "https://ui-avatars.com/api/?name=" + url.QueryEscape(user.Name) + "&background=random"
	}

	return dto.UserData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		AvatarUrl: avatarUrl,
		Role:      "",
	}, nil
}
