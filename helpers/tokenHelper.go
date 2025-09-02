package helpers

import (
	"errors"
	"log"
	"time"

	"github.com/abdallahelassal/go-jwt-authentication-.git/models"
	"github.com/golang-jwt/jwt/v5"
)

// type Claims struct {
//     UserID uint   `json:"user_id"`
//     Email  string `json:"email"`
//     jwt.RegisteredClaims
// }

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}



func GenerateToken(user *models.User, tokenType string) (string, error) {
	var expritionTime time.Time
	var subject string

	switch tokenType {
	case "access":
		expritionTime = time.Now().Add(time.Hour * 1)
	case "refresh":
		expritionTime = time.Now().Add(time.Hour * 2)
		subject = "refresh"
	default:
		log.Printf("token not receive %s", tokenType)
		return "", errors.New("invaled token type")
	}
	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expritionTime),
			Issuer:    "myapp",
			Subject:   subject,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(AppConfig.SECRET_KEY))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}


func ValidateToken(tokenString string, tokenType string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(AppConfig.SECRET_KEY), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return  nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if tokenType == "refresh" && claims.Subject != "refresh" {
		return nil, errors.New("invalid token type")
	}

	return claims, nil

}
