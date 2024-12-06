package impl

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hafidzyami/jaundicebe/model"
	"github.com/hafidzyami/jaundicebe/repository"
	supa "github.com/nedpals/supabase-go"
)

type hospitalRepositoryImpl struct {
	client *supa.Client
}

func NewHospitalRepository(client *supa.Client) repository.HospitalRepository {
	return &hospitalRepositoryImpl{client: client}
}

func (repository *hospitalRepositoryImpl) Create(ctx context.Context, hospital model.HospitalCreateOrUpdate) model.HospitalCreateOrUpdate {
	var results []model.HospitalCreateOrUpdate
	err := repository.client.DB.From("hospitals").Insert(hospital).Execute(&results)
	if err != nil {
		panic(err)
	}
	if len(results) > 0 {
		return results[0] // Return the first inserted hospital
	}
	return model.HospitalCreateOrUpdate{}
}

func (repository *hospitalRepositoryImpl) Update(ctx context.Context, hospital model.Hospital) model.Hospital {
	var results []model.Hospital
	err := repository.client.DB.From("hospitals").
		Update(map[string]interface{}{
			"name":     hospital.Name,
			"city":     hospital.City,
			"province": hospital.Province,
			"image_url":    hospital.ImageURL,
			"contact":  hospital.Contact,
		}).
		Eq("id", strconv.Itoa(hospital.ID)).
		Execute(&results)
	if err != nil {
		panic(err)
	}
	if len(results) > 0 {
		return results[0] // Return the first inserted hospital
	}
	return model.Hospital{}
}

func (repository *hospitalRepositoryImpl) Delete(ctx context.Context, hospital model.Hospital)  {
	err := repository.client.DB.From("hospitals").Delete().Eq("id", strconv.Itoa(hospital.ID)).Execute(nil)
	if err != nil {
		panic(err)
	}
}

func (repository *hospitalRepositoryImpl) FindByID(ctx context.Context, id int) (model.Hospital, error) {
	var results []model.Hospital
	err := repository.client.DB.From("hospitals").Select("*").Eq("id", strconv.Itoa(id)).Execute(&results)
	if err != nil {
		return model.Hospital{}, err
	}
	if len(results) > 0 {
		return results[0], nil
	}
	return model.Hospital{}, fmt.Errorf("hospital with id %d not found", id)
}

func (repository *hospitalRepositoryImpl) FindAll(ctx context.Context) []model.Hospital {
	var results []model.Hospital
	err := repository.client.DB.From("hospitals").Select("*").Execute(&results)
	if err != nil {
		panic(err)
	}
	return results
}
