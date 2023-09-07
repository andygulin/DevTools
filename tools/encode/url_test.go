package encode

import (
	"testing"
)

var obj Url

func init() {
	obj = new(UrlImpl)
}

func TestUrlImpl_Encode(t *testing.T) {
	str := "你好，世界。"
	output, _ := obj.Encode(str)
	t.Log(output)
}

func TestUrlImpl_Decode(t *testing.T) {
	str := "%E4%BD%A0%E5%A5%BD%EF%BC%8C%E4%B8%96%E7%95%8C%E3%80%82"
	output, _ := obj.Decode(str)
	t.Log(output)
}
