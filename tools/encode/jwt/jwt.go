package jwt

import "github.com/golang-jwt/jwt/v4"

type KeyValue map[string]any
type Claims struct {
	KeyValues []KeyValue
	jwt.RegisteredClaims
}
type Jwt interface {
	Create(claims *Claims) (string, error)
	Parse(jwtToken string) (*Claims, error)
}

type JwtImpl struct {
	JwtKey     string
	ExpireTime int64
}

func (obj *JwtImpl) Create(claims *Claims) (string, error) {
	return "", nil
}

func (obj *JwtImpl) Parse(jwtToken string) (*Claims, error) {
	return nil, nil
}
