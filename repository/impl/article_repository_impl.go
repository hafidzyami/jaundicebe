package impl

import (
	"context"
	"fmt"
	"github.com/hafidzyami/jaundicebe/model"
	"github.com/hafidzyami/jaundicebe/repository"
	supa "github.com/nedpals/supabase-go"
	"strconv"
)

type articleRepositoryImpl struct {
	client *supa.Client
}

func NewArticleRepository(client *supa.Client) repository.ArticleRepository {
	return &articleRepositoryImpl{client: client}
}

func (repository *articleRepositoryImpl) Create(ctx context.Context, article model.ArticleCreateOrUpdate) model.ArticleCreateOrUpdate {
	var results []model.ArticleCreateOrUpdate // fixcannot unmarshal array into Go value of type model.Article
	err := repository.client.DB.From("article").Insert(article).Execute(&results)
	if err != nil {
		panic(err)
	}
	if len(results) > 0 {
		return results[0] // Return the first inserted article
	}
	return model.ArticleCreateOrUpdate{}
}

func (repository *articleRepositoryImpl) Update(ctx context.Context, article model.Article) model.Article {
	var results []model.Article
	err := repository.client.DB.From("article").
		Update(map[string]interface{}{
			"title":       article.Title,
			"description": article.Description,
			"link":        article.Link,
		}).
		Eq("id", strconv.Itoa(article.ID)).
		Execute(&results)
	if err != nil {
		panic(err)
	}
	if len(results) > 0 {
		return results[0] // Return the first inserted article
	}
	return model.Article{}
}

func (repository *articleRepositoryImpl) Delete(ctx context.Context, article model.Article) {
	err := repository.client.DB.From("article").Delete().Eq("id", strconv.Itoa(article.ID)).Execute(nil)
	if err != nil {
		panic(err)
	}
}

func (repository *articleRepositoryImpl) FindByID(ctx context.Context, id int) (model.Article, error) {
	var results []model.Article
	err := repository.client.DB.From("article").Select("*").Eq("id", strconv.Itoa(id)).Execute(&results)
	if err != nil {
		return model.Article{}, err
	}
	if len(results) == 0 {
		return model.Article{}, fmt.Errorf("article with id %d not found", id)
	}
	return results[0], nil // Return the first result
}

func (repository *articleRepositoryImpl) FindAll(ctx context.Context) []model.Article {
	var results []model.Article
	err := repository.client.DB.From("article").Select("*").Execute(&results)
	if err != nil {
		panic(err)
	}
	return results
}
