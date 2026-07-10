package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}
	return match, nil
}

func GetToken(req http.Request) (string, error) {
	rawTokenString := req.Header.Get("Authorization")
	if rawTokenString == "" {
		return "", fmt.Errorf("Authorization header missing")
	}
	token := strings.TrimPrefix(rawTokenString, "Bearer ")
	return token, nil
}
