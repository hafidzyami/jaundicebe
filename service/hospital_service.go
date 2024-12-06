package service

import (
	"context"
	"github.com/hafidzyami/jaundicebe/model"
)

type HospitalService interface {
	Create(ctx context.Context, model model.HospitalCreateOrUpdate) model.HospitalCreateOrUpdate
	Update(ctx context.Context, model model.HospitalCreateOrUpdate, id int) model.HospitalCreateOrUpdate
	Delete(ctx context.Context, id int)
	FindByID(ctx context.Context, id int) (model.Hospital, error)
	FindAll(ctx context.Context) []model.Hospital
}