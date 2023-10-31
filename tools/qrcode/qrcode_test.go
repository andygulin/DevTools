package qrcode

import "testing"

var obj QrCode

func TestQrCodeImpl_Create(t *testing.T) {
	obj = &QrCodeImpl{Size: 500}
	output, _ := obj.Create("https://www.baidu.com")
	t.Log(output)
}
