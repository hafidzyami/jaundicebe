package repository

import (
	"context"
	"github.com/hafidzyami/jaundicebe/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.UserCreateOrUpdate) (model.UserCreateOrUpdate, error)
	UpdatePassword(ctx context.Context, user model.User) (model.User, error)
	FindByUsername(ctx context.Context, username string) (model.User, error)
	FindByID(ctx context.Context, id string) (model.User, error)
}