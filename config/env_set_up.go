package config

import (
	"log"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// to load env variables
func loadEnv() {
	//why over load? // directly loading the file
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatal("error loading .enc file", err)
	}
}

func GetEnvWithKey(key string, defaultValue string) string {
	//syscall
	keyVal, found := syscall.Getenv(key)
	if !found {
		syscall.Setenv(key, defaultValue)
		return defaultValue
	}
	return keyVal

}

func init() {
	loadEnv()

	if GetEnvWithKey("APP_ENVIRONMENT", "dev") == "pro" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

}
