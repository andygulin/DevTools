package passwd

import (
	"math/rand"
	"time"
)

type Password interface {
	RandomPassword() (string, error)
}

type PasswordImpl struct {
	Length int
}

var str = []rune("abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789!@#$%^&*()_+{}[]")

func (obj *PasswordImpl) RandomPassword() (string, error) {
	rand.New(rand.NewSource(time.Now().Unix()))
	pass := make([]rune, obj.Length)
	for i := 0; i < obj.Length; i++ {
		pass[i] = str[rand.Intn(len(str))]
	}
	return string(pass), nil
}
