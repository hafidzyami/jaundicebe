package model

type Hospital struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	City     string `json:"city"`
	Province string `json:"province"`
	ImageURL string `json:"image_url"`
	Contact  string `json:"contact"`
}

type HospitalCreateOrUpdate struct {
	Name     string `json:"name" validate:"required"`
	City     string `json:"city" validate:"required"`
	Province string `json:"province" validate:"required"`
	ImageURL string `json:"image_url" validate:"required,url"`
	Contact  string `json:"contact" validate:"required"`
}
