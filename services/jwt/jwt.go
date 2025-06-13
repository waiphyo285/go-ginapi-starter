package jwtservice

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"neohub.asia/mod/config"
)

func CreateToken(data map[string]interface{}, expiresDelta time.Duration) (string, error) {
	claims := jwt.MapClaims{}
	for k, v := range data {
		claims[k] = v
	}
	claims["exp"] = time.Now().Add(expiresDelta).Unix()

	token := jwt.NewWithClaims(jwt.GetSigningMethod(config.JWT_ALGO), claims)
	return token.SignedString(config.JWT_SECRET)
}

func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return config.JWT_SECRET, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}
