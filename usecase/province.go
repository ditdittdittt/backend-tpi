package usecase

import (
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
		return nil, err
	}

	return provinces, nil
}

func NewProvinceUsecase(provinceClientUsecase client.ProvinceRepository) ProvinceUsecase {
	return &provinceUsecase{provinceClientUsecase: provinceClientUsecase}
}
