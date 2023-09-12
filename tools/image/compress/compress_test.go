package compress

import "testing"

func TestCompressImpl_Compress(t *testing.T) {
	obj := new(CompressImpl)
	obj.Quality = 50
	output, err := obj.Compress("sample.jpg")
	if err != nil {
		t.Log(err)
	}
	t.Log(output)
}
