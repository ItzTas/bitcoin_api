package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	dat, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func NewJWT(dbu database.User, tokenSecret string, expiresIn time.Duration) (string, error) {
	signingKey := []byte(tokenSecret)
	id := dbu.ID.String()
	if id == "" {
		return "", errors.New("invalid id")
	}
	claims := jwt.RegisteredClaims{
		Issuer:    "bitcoin_api",
		Subject:   id,
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func GetBearerToken(header http.Header) (string, error) {
	token := header.Get("Authorization")
	if token == "" {
		return "", errors.New("empty auth header")
	}

	slides := strings.Split(token, " ")
	if len(slides) != 2 || slides[0] != "Bearer" {
		return "", errors.New("bad formatted auth header")
	}

	return slides[1], nil
}

func GetIDByToken(token, secretKey string) (string, error) {
	tok, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil {
		return "", fmt.Errorf("could not get token error: \n%v", err)
	}

	if !tok.Valid {
		return "", errors.New("invalid token")
	}
	id, err := tok.Claims.GetSubject()
	return id, err
}
