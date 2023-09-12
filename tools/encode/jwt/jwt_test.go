package jwt

import "testing"

var obj Jwt

var jwtKey = "u*29HJ1320"

func init() {
	obj = &JwtImpl{
		JwtKey:       jwtKey,
		ExpireMinute: 30,
	}
}

func TestJwtImpl_Create(t *testing.T) {
	keyValue := make(map[string]any)
	keyValue["userId"] = 1
	token, err := obj.Create(1)
	if err != nil {
		t.Log(err)
	}
	t.Log(token)
}

func TestJwtImpl_Parse(t *testing.T) {
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTY5NDUxNzQ4MH0.2vethnAEuwhj0CmhHpaLSENNUHwd-7t18RPKGNrwwGQ"
	userId, err := obj.Parse(jwtToken)
	if err != nil {
		t.Log(err)
	}
	t.Log(userId)
}
