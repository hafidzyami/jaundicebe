package service

import(
	"context"
	"github.com/hafidzyami/jaundicebe/model"
)

type UserService interface {
	Create(ctx context.Context, user model.UserCreateOrUpdate) (model.UserCreateOrUpdate, error)
	Login(ctx context.Context, user model.UserCreateOrUpdate) (string, error)
	ChangePassword(ctx context.Context, userId string, passwords model.ChangePassword) (string, error)
}