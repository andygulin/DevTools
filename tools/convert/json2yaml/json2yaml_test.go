package json2yaml

import "testing"

func TestJson2YamlImpl_Convert(t *testing.T) {
	obj := new(Json2YamlImpl)
	output, err := obj.Convert("sample.json")
	if err != nil {
		t.Log(err)
	}
	t.Log(output)
}
