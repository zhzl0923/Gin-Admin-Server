package app

import (
	"gin-admin/global"
	"time"

	"github.com/golang-jwt/jwt"
)

type AdminClaims struct {
	UserID uint64
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(userID uint64) (string, error) {
	expireTime := time.Now().Add(global.JWTSetting.Expire)
	claims := &AdminClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(GetJWTSecret())
}

func ParseToken(token string) (*AdminClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &AdminClaims{}, func(t *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*AdminClaims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
