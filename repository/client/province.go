package client

import (
	"encoding/json"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
)

type ProvinceRepository interface {
	Get() ([]entities.Province, error)
}

type provinceRepository struct {
}

func (p *provinceRepository) Get() ([]entities.Province, error) {
	var response GetProvinceResponse
	result := make([]entities.Province, 0)

	path := "https://dev.farizdotid.com/api/daerahindonesia/provinsi"

	responseByte, err := helper.CallEndpoint(map[string]string{}, path, "GET")
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(responseByte, &response)

	for _, provinsi := range response.Provinsi {
		province := entities.Province{
			ID:   provinsi.ID,
			Name: provinsi.Nama,
		}
		result = append(result, province)
	}

	return result, nil
}

func NewProvinceRepository() ProvinceRepository {
	return &provinceRepository{}
}
