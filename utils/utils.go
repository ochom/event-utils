package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"regexp"
)

// GetEnvOrDefault get env variable or return default value
func GetEnvOrDefault(key, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}

// MustGetEnv gets env variable or return error
func MustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("Environment variable %s not set", key))
	}

	return val
}

// GenerateOTP create a length size string of numbers for OTP
func GenerateOTP(size int) string {
	seed := "0123456789"
	otp := make([]byte, size)

	for i := 0; i < size; i++ {
		max := big.NewInt(int64(len(seed)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "1234"
		}

		otp[i] = seed[num.Int64()]
	}
	return string(otp)
}

// RenameFile renames a file keeping the actual file extension
func RenameFile(oldName, newName string) string {
	ext := regexp.MustCompile(`\.([a-zA-Z]+)$`).FindStringSubmatch(oldName)[1]
	re := regexp.MustCompile(`(.*)\.([a-zA-Z]+)$`)
	return re.ReplaceAllString(oldName, fmt.Sprintf("%s.%s", newName, ext))
}

// ParseMobile ...
func ParseMobile(phone string) (string, error) {

	// replace all non digits with empty string
	reg := regexp.MustCompile("[^0-9]+")
	phone = reg.ReplaceAllString(phone, "")

	rex := regexp.MustCompile(`^0`)
	phone = rex.ReplaceAllString(phone, "254")

	// if length is 9 add 254
	if len(phone) == 9 {
		phone = "254" + phone
	}

	if len(phone) != 12 {
		return "", fmt.Errorf("invalid phone number")
	}

	return phone, nil
}
