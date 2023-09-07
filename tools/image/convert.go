package image

type Convert interface {
	Convert(input string, output string) error
}

type ConvertImpl struct {
}

func (obj *ConvertImpl) Convert(input string, output string) error {
	return nil
}
