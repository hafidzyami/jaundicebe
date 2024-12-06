package service

import (
	"context"
	"github.com/hafidzyami/jaundicebe/model"
)

type ArticleService interface {
	Create(ctx context.Context, model model.ArticleCreateOrUpdate) model.ArticleCreateOrUpdate
	Update(ctx context.Context, model model.ArticleCreateOrUpdate, id int) model.ArticleCreateOrUpdate
	Delete(ctx context.Context, id int)
	FindByID(ctx context.Context, id int) (model.Article, error)
	FindAll(ctx context.Context) []model.Article
}