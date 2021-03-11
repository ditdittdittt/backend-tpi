package client

type GetDistrictByProvinceIDResponse struct {
	KotaKabupaten []struct {
		ID         int    `json:"id"`
		IDProvinsi string `json:"id_provinsi"`
		Nama       string `json:"nama"`
	} `json:"kota_kabupaten"`
}

type GetProvinceResponse struct {
	Provinsi []struct {
		ID   int    `json:"id"`
		Nama string `json:"nama"`
	} `json:"provinsi"`
}
