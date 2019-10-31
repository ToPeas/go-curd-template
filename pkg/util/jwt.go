package util

import (
	"github.com/dgrijalva/jwt-go"
	"github/ToPeas/go-curd-templatepkg/setting"
)

var jwtSecret = []byte(setting.Config.App.JwtSecret)

type Claims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

func GenerateToken(uid int64) (string, error) {
	claims := Claims{
		uid,
		jwt.StandardClaims{},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
