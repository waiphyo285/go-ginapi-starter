package jwtservice

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"neohub.asia/mod/config"
)

func CreateToken(data map[string]interface{}) (string, error) {

	claims := jwt.MapClaims{}
	jwtCfg := config.LoadJWTConfig()

	for k, v := range data {
		claims[k] = v
	}
	now := time.Now()
	claims["iat"] = now.Unix() 
	claims["exp"] = now.Add(jwtCfg.ExpiresIn).Unix() 

	token := jwt.NewWithClaims(jwt.GetSigningMethod(jwtCfg.Algorithm), claims)
	return token.SignedString(jwtCfg.Secret)
}

func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	jwtCfg := config.LoadJWTConfig()
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtCfg.Secret, nil
	})


	if err != nil {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}
