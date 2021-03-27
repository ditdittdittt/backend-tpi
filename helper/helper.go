package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"
	"golang.org/x/crypto/bcrypt"

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

func HashAndSaltPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePassword(hashedPwd string, pwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func GetCurrentUserID(c *gin.Context) (int, int, int, error) {
	userID := 0
	tpiID := 0
	districtID := 0

	curUserID, ok := c.Get("userID")
	if !ok {
		return userID, tpiID, districtID, stacktrace.NewError("Invalid user")
	}
	userID = curUserID.(int)

	curTpiID, ok := c.Get("tpiID")
	if ok {
		tpiID = curTpiID.(int)
	}

	curDistrictID, ok := c.Get("districtID")
	if ok {
		districtID = curDistrictID.(int)
	}

	return userID, tpiID, districtID, nil
}
