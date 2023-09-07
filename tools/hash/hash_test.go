package hash

import "testing"

var obj Hash

const str = "Hello World~"

func init() {
	obj = new(HashImpl)
}

func TestHashImpl_MD5(t *testing.T) {
	ret, _ := obj.MD5(str)
	t.Log(ret)
}

func TestHashImpl_SHA1(t *testing.T) {
	ret, _ := obj.SHA1(str)
	t.Log(ret)
}

func TestHashImpl_SHA256(t *testing.T) {
	ret, _ := obj.SHA256(str)
	t.Log(ret)
}

func TestHashImpl_SHA512(t *testing.T) {
	ret, _ := obj.SHA512(str)
	t.Log(ret)
}
