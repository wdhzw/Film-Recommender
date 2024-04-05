package movie_server

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Mysql *gorm.DB
)

func InitMysql() {
	dsn := "guohaonan:ghn980421@tcp(cs5224-movie_server-meta.c3o8o8eyqv2b.us-east-1.rds.amazonaws.com:3306)/movie?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect to movie_server database, err:%s\n", err))
	}

}