package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/ditdittdittt/backend-tpi/entities"
)

func ValidatePermission(permissionList []*entities.Permission, permissionNeeded string) bool {
	for _, permission := range permissionList {
		if permissionNeeded == permission.Name {
			return true
		}
	}
	return false
}

func CallEndpoint(requestBody map[string]string, path string, method string) ([]byte, error) {
	jsonRequest, _ := json.Marshal(requestBody)

	var payload = bytes.NewReader(jsonRequest)

	request, err := http.NewRequest(method, path, payload)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// check for valid json response
	var js map[string]interface{}
	err = json.Unmarshal(responseByte, &js)
	if err != nil {
		return nil, err
	}

	return responseByte, nil
}

func MakeExcel() {
	xlsx := excelize.NewFile()
	sheet1Name := "Sheet One"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Name")
	xlsx.SetCellValue(sheet1Name, "B1", "Gender")
	xlsx.SetCellValue(sheet1Name, "C1", "Age")

	err := xlsx.AutoFilter(sheet1Name, "A1", "C1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

}
