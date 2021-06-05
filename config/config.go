package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

const (
	dev  = "development"
	prod = "production"
)

var (
	Dsn            string
	TimeoutContext time.Duration
	JwtSecret      []byte
)

func Init() {
	if getAppEnv() == prod {
		initProd()
	} else {
		initDev()
	}
}

func getAppEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return dev
	}

	return prod
}

func initDev() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	//val := url.Values{}
	//val.Add("parseTime", "1")
	//val.Add("loc", "Asia/Jakarta")
	//Dsn = fmt.Sprintf("%s?%s", connection, val.Encode())

	TimeoutContext = time.Duration(viper.GetInt("context.timeout")) * time.Second

	JwtSecret = []byte(viper.GetString(`jwt_secret`))
}

func initProd() {
	dbHost := os.Getenv(`DB_HOST`)
	dbPort := os.Getenv(`DB_PORT`)
	dbUser := os.Getenv(`DB_USER`)
	dbPass := os.Getenv(`DB_PASS`)
	dbName := os.Getenv(`DB_NAME`)

	Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	timeout, _ := strconv.Atoi(os.Getenv("TIMEOUT"))
	TimeoutContext = time.Duration(timeout) * time.Second

	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
}
