package xml

import "testing"

func TestXmlFormat_Format(t *testing.T) {
	obj := new(XmlFormat)
	input := "<name>aaa</name><age>11</age>"
	output, err := obj.Format(input)
	if err != nil {
		return
	}
	t.Log(output)
}
