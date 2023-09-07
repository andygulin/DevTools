package json2yaml

type Json2Yaml interface {
	Convert(input string) (string, error)
}

type Json2YamlImpl struct {
}

func (obj *Json2YamlImpl) Convert(input string) (string, error) {

	return "", nil
}
