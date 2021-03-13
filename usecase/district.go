package usecase

import (
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type DistrictUsecase interface {
	Index(provinceID int) ([]entities.District, error)
}

type districtUsecase struct {
	districtRepository mysql.DistrictRepository
}

func (d *districtUsecase) Index(provinceID int) ([]entities.District, error) {
	query := map[string]interface{}{
		"province_id": provinceID,
	}

	districts, err := d.districtRepository.Get(query)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Get] District repository error")
	}

	return districts, nil
}

func NewDistrictUsecase(districtRepository mysql.DistrictRepository) DistrictUsecase {
	return &districtUsecase{districtRepository: districtRepository}
}
