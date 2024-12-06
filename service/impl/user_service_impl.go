package impl

import (
	"context"
	"fmt"

	"github.com/hafidzyami/jaundicebe/model"
	"github.com/hafidzyami/jaundicebe/repository"
	"github.com/hafidzyami/jaundicebe/service"
	"github.com/hafidzyami/jaundicebe/utils"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{
		userRepository: *userRepository,
	}
}

func (s *userServiceImpl) Create(ctx context.Context, user model.UserCreateOrUpdate) (model.UserCreateOrUpdate, error) {
	return s.userRepository.Create(ctx, user)
}

func (s *userServiceImpl) ChangePassword(ctx context.Context, userId string, passwords model.ChangePassword) (string, error) {
	getUser, err := s.userRepository.FindByID(ctx, userId)
	if err != nil {
		return "", fmt.Errorf("invalid user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(passwords.OldPassword))
	if err != nil {
		return "", fmt.Errorf("invalid old password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwords.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password")
	}

	getUser.Password = string(hashedPassword)
	_, err = s.userRepository.UpdatePassword(ctx, getUser)
	if err != nil {
		return "", fmt.Errorf("error updating password")
	}

	return "password updated", nil
}


func (s *userServiceImpl) Login(ctx context.Context, user model.UserCreateOrUpdate) (string, error) {
	getUser, err := s.userRepository.FindByUsername(ctx, user.Username)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateToken(getUser.ID)
	if err != nil {
		return "", fmt.Errorf("error generating token")
	}

	return token, nil
}