package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
)

type DistrictRepository interface {
	GetByProvinceID(id int) ([]entities.District, error)
}

type districtRepository struct {
}

func (d *districtRepository) GetByProvinceID(id int) ([]entities.District, error) {
	var response GetDistrictByProvinceIDResponse
	result := make([]entities.District, 0)

	path := fmt.Sprintf("https://dev.farizdotid.com/api/daerahindonesia/kota?id_provinsi=%d", id)

	responseByte, err := helper.CallEndpoint(map[string]string{}, path, "GET")
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(responseByte, &response)

	for _, kotaKabupaten := range response.KotaKabupaten {
		provinceID, _ := strconv.Atoi(kotaKabupaten.IDProvinsi)
		district := entities.District{
			ID:         kotaKabupaten.ID,
			ProvinceID: provinceID,
			Name:       kotaKabupaten.Nama,
		}
		result = append(result, district)
	}

	return result, nil
}

func NewDistrictRepository() DistrictRepository {
	return &districtRepository{}
}
