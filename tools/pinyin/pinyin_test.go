package pinyin

import "testing"

var obj Pinyin

func init() {
	obj = new(PinyinImpl)
}

func TestPinyinImpl_Convert(t *testing.T) {
	output, _ := obj.Convert("你好")
	t.Log(output)
}
