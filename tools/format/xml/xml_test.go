package xml

import "testing"

func TestXmlFormat_Format(t *testing.T) {
	obj := new(XmlFormat)
	input := "<user><name>aaa</name><age>11</age></user>"
	output, err := obj.Format(input)
	if err != nil {
		return
	}
	t.Log(output)
}

func TestXmlFormat_FormatFile(t *testing.T) {
	obj := new(XmlFormat)
	output, err := obj.FormatFile("sample.xml")
	if err != nil {
		return
	}
	t.Log(output)
}
