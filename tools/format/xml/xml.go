package xml

import "github.com/go-xmlfmt/xmlfmt"

type XmlFormat struct {
}

func (obj *XmlFormat) Format(input string) (string, error) {
	return xmlfmt.FormatXML(input, "", "	", true), nil
}
