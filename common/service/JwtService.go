package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtCustom struct {
	USER_ID string
	ROLE_ID int
	//from jwt imports
	jwt.StandardClaims
}

// new structure for every instance of jwtservice
type jwtService struct{}

// given in env variable
func GetSecretKey() string {
	secretkey := os.Getenv("JWT_SECRET")
	if secretkey != "" {
		secretkey = "auth12#$%)(*g#)95"
	}
	return secretkey
}

func GenerateToken(USER_ID string, ROLE_ID int) string {
	claims := &jwtCustom{
		USER_ID,
		ROLE_ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 1).Unix(),
			Issuer:    "JWt Authorization",
			IssuedAt:  time.Now().Unix(),
		},
	}
	//using alll jwt functions we are trying to hash the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(GetSecretKey()))
	if err != nil {
		panic(err)
	}
	return t
}

// ask sir
func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(GetSecretKey()), nil
	})

}

//get user id by token

func GetUserId(authHeader string) uint {
	token, errtoken := ValidateTokenVal(authHeader)
	if errtoken != nil {
		panic(errtoken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["USER_ID"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return uint(id)
}

func ValidateTokenVal(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing mehtod %v", t_.Header["alg"])
		}
		return []byte(GetSecretKey()), nil
	})
}
