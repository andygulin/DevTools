package yaml2json

type Yaml2Json interface {
	Convert(input string) (string, error)
}

type Yaml2JsonImpl struct {
}

func (obj *Yaml2JsonImpl) Convert(input string) (string, error) {

	return "", nil
}
