package uuid

import (
	"github.com/google/uuid"
)

type UUID interface {
	GenerateUUID() (string, error)
}

type GoogleUUID struct {
}

func (obj *GoogleUUID) GenerateUUID() (string, error) {
	return uuid.New().String(), nil
}
