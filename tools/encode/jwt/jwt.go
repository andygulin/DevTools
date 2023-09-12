package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Jwt interface {
	Create(userId int) (string, error)
	Parse(jwtToken string) (int, error)
}

type JwtImpl struct {
	JwtKey       string
	ExpireMinute int // 分钟
}

type Claims struct {
	UserId int
	jwt.RegisteredClaims
}

func (obj *JwtImpl) Create(userId int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Minute * time.Duration(obj.ExpireMinute))
	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireTime},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(obj.JwtKey)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (obj *JwtImpl) Parse(jwtToken string) (int, error) {
	tokenClaims, _ := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(obj.JwtKey), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			fmt.Println()
			return claims.UserId, nil
		}
	}
	return 0, nil
}
