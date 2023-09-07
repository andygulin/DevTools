package base64

import (
	"encoding/base64"
	"os"
)

type TypeText string
type TypeImageFile string

type Base64 interface {
	EncodeText(input TypeText) (string, error)
	EncodeImage(input TypeImageFile) (string, error)
}

type Base64Impl struct {
}

func (obj *Base64Impl) EncodeText(input TypeText) (string, error) {
	bytes := []byte(input)
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func (obj *Base64Impl) EncodeImage(input TypeImageFile) (string, error) {
	imageData, err := os.ReadFile(string(input))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imageData), nil
}
