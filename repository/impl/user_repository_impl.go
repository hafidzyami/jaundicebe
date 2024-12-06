package impl

import (
	"context"
	"github.com/hafidzyami/jaundicebe/model"
	"github.com/hafidzyami/jaundicebe/repository"
	supa "github.com/nedpals/supabase-go"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"strings"
)

type userRepositoryImpl struct {
	client *supa.Client
}

func NewUserRepository(client *supa.Client) repository.UserRepository {
	return &userRepositoryImpl{client: client}
}

func (repository *userRepositoryImpl) Create(ctx context.Context, user model.UserCreateOrUpdate) (model.UserCreateOrUpdate, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)

	var results []model.UserCreateOrUpdate
	err = repository.client.DB.From("users").Insert(user).Execute(&results)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") { // Adjust for your DB's error message
			return model.UserCreateOrUpdate{}, fmt.Errorf("username already exists")
		}
		return model.UserCreateOrUpdate{}, err
	}
	if len(results) > 0 {
		return results[0], nil // Return the first inserted user
	}
	return model.UserCreateOrUpdate{}, nil
}

func (repository *userRepositoryImpl) UpdatePassword(ctx context.Context, user model.User) (model.User, error) {
	var results []model.User
	err := repository.client.DB.From("users").
		Update(map[string]interface{}{
			"password": user.Password,
		}).
		Eq("username", user.Username).
		Execute(&results)
	if err != nil {
		return model.User{}, err
	}
	if len(results) > 0 {
		return results[0], nil // Return the first inserted user
	}
	return model.User{}, nil
}

func (repository *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (model.User, error) {
	var results []model.User
	err := repository.client.DB.From("users").Select("*").Eq("username", username).Execute(&results)
	if err != nil {
		return model.User{}, err
	}
	if len(results) == 0 {
		return model.User{}, fmt.Errorf("user with username %s not found", username)
	}
	return results[0], nil
}

func (repository *userRepositoryImpl) FindByID(ctx context.Context, id string) (model.User, error) {
	var results []model.User
	err := repository.client.DB.From("users").Select("*").Eq("id", id).Execute(&results)
	if err != nil {
		return model.User{}, err
	}
	if len(results) == 0 {
		return model.User{}, fmt.Errorf("user with id %s not found", id)
	}
	return results[0], nil
}
