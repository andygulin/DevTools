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
	token, err := obj.Create(30227)
	if err != nil {
		t.Log(err)
	}
	t.Log(token)
}

func TestJwtImpl_Parse(t *testing.T) {
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMwMjI3LCJleHAiOjE2OTY2NjY1NTB9.oFxTjzJhk6Px9Lka3JR0jiCvAGwKBMENWzdbTKbHswU"
	userId, err := obj.Parse(jwtToken)
	if err != nil {
		t.Log(err)
	}
	t.Log(userId)
}
