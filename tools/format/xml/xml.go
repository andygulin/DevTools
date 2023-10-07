package xml

import (
	"github.com/go-xmlfmt/xmlfmt"
	"os"
	"path/filepath"
)

type XmlFormat struct {
}

func (obj *XmlFormat) Format(input string) (string, error) {
	return xmlfmt.FormatXML(input, "", "	", true), nil
}

func (obj *XmlFormat) FormatFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	str, err := obj.Format(string(content))
	if err != nil {
		return "", err
	}
	err = os.WriteFile(filename, []byte(str), 0644)
	if err != nil {
		return "", err
	}

	output, _ := filepath.Abs(filename)
	return output, nil
}
