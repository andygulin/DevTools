package qrcode

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"path/filepath"
)

type QrCode interface {
	Create(input string) (string, error)
}

type QrCodeImpl struct {
	Size int // px
}

func (obj *QrCodeImpl) Create(input string) (string, error) {
	fileName := fmt.Sprintf("qr_code_%d.png", obj.Size)
	err := qrcode.WriteFile(input, qrcode.Medium, obj.Size, fileName)
	output, _ := filepath.Abs(fileName)
	if err != nil {
		return "", err
	}
	return output, nil
}
