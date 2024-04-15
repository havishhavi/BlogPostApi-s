package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connects to the sql
func GoConnect() *gorm.DB {
	//godotenv is a go library
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load env files")
	}
	mySQLHost := os.Getenv("DB_HOST")
	mySQLUser := os.Getenv("DB_USER")
	mySQLPass := os.Getenv("DB_PASS")
	mySQLDBName := os.Getenv("DB_NAME")
	mySQLDBPort := os.Getenv("DB_PORT")

	//dsn := "root:admin@tcp(127.0.0.1:3306)/blogex?charset=utf8mb4"
	//get the above variables from env and load them in the above form
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mySQLUser, mySQLPass, mySQLHost, mySQLDBPort, mySQLDBName)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("unable to create connection with mysql" + err.Error())
	}
	return db

}
