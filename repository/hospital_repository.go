package repository

import (
	"context"
	"github.com/hafidzyami/jaundicebe/model"
)

type HospitalRepository interface {
	Create(ctx context.Context, hospital model.HospitalCreateOrUpdate) model.HospitalCreateOrUpdate
	Update(ctx context.Context, hospital model.Hospital) model.Hospital
	Delete(ctx context.Context, hospital model.Hospital)
	FindByID(ctx context.Context, id int) (model.Hospital, error)
	FindAll(ctx context.Context) []model.Hospital
}