package base64

import "testing"

var obj Base64

func init() {
	obj = new(Base64Impl)
}

func TestBase64Impl_EncodeText(t *testing.T) {
	input := "Hello World~"
	t.Log(obj.EncodeText(TypeText(input)))
}

func TestBase64Impl_EncodeImage(t *testing.T) {
	input := "image.jpeg"
	t.Log(obj.EncodeImage(TypeImageFile(input)))
}
