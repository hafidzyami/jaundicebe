package model

type Article struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type ArticleCreateOrUpdate struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Link		string `json:"link" validate:"required,url"`
}