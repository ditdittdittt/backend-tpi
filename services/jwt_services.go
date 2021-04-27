package services

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type JWTService interface {
	GenerateToken(user *entities.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	ExtractClaims(tokenStr string) (string, error)
}
type AuthCustomClaims struct {
	Username string `json:"username"`
	RoleID   int    `json:"role_id"`
	Nik      string `json:"nik"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func NewJWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "tpi-app",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(user *entities.User) (string, error) {
	claims := &AuthCustomClaims{
		user.Username,
		user.RoleID,
		user.Nik,
		user.Name,
		user.Address,
		jwt.StandardClaims{
			Id:        strconv.Itoa(user.ID),
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}

func (service *jwtServices) ExtractClaims(tokenStr string) (string, error) {
	hmacSecretString := service.secretKey
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		accessUsername, ok := claims["username"]
		if !ok {
			return "", err
		}
		return accessUsername.(string), nil
	} else {
		log.Printf("Invalid JWT Token")
		return "", err
	}
}
