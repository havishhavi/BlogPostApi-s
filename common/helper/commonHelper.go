package helper

import (
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

/*
This function will trim white space from
structure
*/
func Trimmer(structure interface{}) {
	msValuePtr := reflect.ValueOf(structure).Elem()

	for i := 0; i < msValuePtr.NumField(); i++ {
		field := msValuePtr.Field(i)

		if field.Kind() == reflect.Struct {
			// Recursively call trimmer for nested structs.
			Trimmer(field.Addr().Interface())
		}

		// Ignore fields that don't have the same type as a string
		if field.Type() != reflect.TypeOf("") {
			continue
		}

		str := field.Interface().(string)
		str = strings.TrimSpace(str)
		field.SetString(str)
	}
}

/*
Developer: mudit
Purmpse: convert string to in
*/
func ConvertStoI(s string) (int, error) {
	number, err := strconv.Atoi(s)
	if err != nil {
		//ELog.Error(err.Error())
		return 0, err
	}
	return number, err
}

/*
/*
Developer: mudit
Purmpse: convert int to string
*/
func ConvertItoS(s int) string {
	stringVal := strconv.Itoa(s)
	return stringVal
}

func PwdEncryption(password string) (string, error) {
	pass_byte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(pass_byte), err

}
