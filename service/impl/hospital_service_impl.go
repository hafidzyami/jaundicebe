package impl

import (
	"context"
	"github.com/hafidzyami/jaundicebe/model"
	"github.com/hafidzyami/jaundicebe/repository"
	"github.com/hafidzyami/jaundicebe/service"
)

type hospitalServiceImpl struct {
	repository repository.HospitalRepository
}

func NewHospitalService(repository *repository.HospitalRepository) service.HospitalService {
	return &hospitalServiceImpl{repository: *repository}
}

func (s *hospitalServiceImpl) Create(ctx context.Context, model model.HospitalCreateOrUpdate) model.HospitalCreateOrUpdate {
	s.repository.Create(ctx, model)
	return model
}

func (s *hospitalServiceImpl) Update(ctx context.Context, modelHospital model.HospitalCreateOrUpdate, id int) model.HospitalCreateOrUpdate {
	hospital := model.Hospital{
		ID : id,
		Name : modelHospital.Name,
		City: modelHospital.City,
		Province: modelHospital.Province,
		Contact: modelHospital.Contact,
		ImageURL: modelHospital.ImageURL,
	}
	s.repository.Update(ctx, hospital)
	return modelHospital
}

func (s *hospitalServiceImpl) Delete(ctx context.Context, id int) {
	hospital, err := s.repository.FindByID(ctx, id)
	if err != nil {
		panic(err)
	}
	s.repository.Delete(ctx, hospital)
}

func (s *hospitalServiceImpl) FindByID(ctx context.Context, id int) (model.Hospital, error) {
	hospital, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return model.Hospital{}, err
	}
	return hospital, nil
}

func (s *hospitalServiceImpl) FindAll(ctx context.Context) []model.Hospital {
	hospitals := s.repository.FindAll(ctx)
	return hospitals
}

