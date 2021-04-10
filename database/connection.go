package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/config"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", config.Dsn)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
