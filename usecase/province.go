package usecase

import (
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/client"
)

type ProvinceUsecase interface {
	Index() ([]entities.Province, error)
}

type provinceUsecase struct {
	provinceClientUsecase client.ProvinceRepository
}

func (p *provinceUsecase) Index() ([]entities.Province, error) {
	provinces, err := p.provinceClientUsecase.Get()
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Get] Province client usecase error")
	}

	return provinces, nil
}

func NewProvinceUsecase(provinceClientUsecase client.ProvinceRepository) ProvinceUsecase {
	return &provinceUsecase{provinceClientUsecase: provinceClientUsecase}
}
