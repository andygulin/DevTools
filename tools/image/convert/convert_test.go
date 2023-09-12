package convert

import "testing"

func TestConvertImpl_Convert(t *testing.T) {
	obj := new(ConvertImpl)
	output, err := obj.Convert("sample.jpg")
	if err != nil {
		t.Log(err)
	}
	t.Log(output)
}
