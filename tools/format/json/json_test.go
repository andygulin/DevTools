package json

import "testing"

func TestJsonFormat_Format(t *testing.T) {
	obj := new(JsonFormat)
	input := "{\"name\":\"aa\",\"age\":11}"
	output, err := obj.Format(input)
	if err != nil {
		return
	}
	t.Log(output)
}

func TestJsonFormat_FormatFile(t *testing.T) {

}
