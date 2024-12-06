package repository

import (
	"context"
	"github.com/hafidzyami/jaundicebe/model"
)

type ArticleRepository interface {
	Create(ctx context.Context, article model.ArticleCreateOrUpdate) model.ArticleCreateOrUpdate
	Update(ctx context.Context, article model.Article) model.Article
	Delete(ctx context.Context, article model.Article)
	FindByID(ctx context.Context, id int) (model.Article, error)
	FindAll(ctx context.Context) []model.Article
}