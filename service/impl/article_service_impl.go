package impl

import (
	"context"
	"github.com/hafidzyami/jaundicebe/model"
	"github.com/hafidzyami/jaundicebe/repository"
	"github.com/hafidzyami/jaundicebe/service"
)

type articleServiceImpl struct {
	repository repository.ArticleRepository
}

func NewArticleService(repository *repository.ArticleRepository) service.ArticleService {
	return &articleServiceImpl{repository: *repository}
}

func (service *articleServiceImpl) Create(ctx context.Context, modelArticle model.ArticleCreateOrUpdate) model.ArticleCreateOrUpdate {
	service.repository.Create(ctx, modelArticle)
	return modelArticle
}

func (service *articleServiceImpl) Update(ctx context.Context, modelArticle model.ArticleCreateOrUpdate, id int) model.ArticleCreateOrUpdate {
	article := model.Article{
		ID: 		id,
		Title:       modelArticle.Title,
		Description: modelArticle.Description,
		Link:     modelArticle.Link,
	}
	service.repository.Update(ctx, article)
	return modelArticle
}

func (service *articleServiceImpl) Delete(ctx context.Context, id int) {
	article, err := service.repository.FindByID(ctx, id)
	if err != nil {
		panic(err)
	}
	service.repository.Delete(ctx, article)
}

func (service *articleServiceImpl) FindByID(ctx context.Context, id int) (model.Article, error) {
	article, err := service.repository.FindByID(ctx, id)
	if err != nil {
		return model.Article{}, err
	}
	return article, nil
}

func (service *articleServiceImpl) FindAll(ctx context.Context) (responses []model.Article) {
	article := service.repository.FindAll(ctx)
	for _, value := range article {
		responses = append(responses, model.Article{
			ID:          value.ID,
			Title:       value.Title,
			Description: value.Description,
			Link:     value.Link,
		})
	}
	if len(responses) == 0 {
		return []model.Article{}
	}
	return responses
}
