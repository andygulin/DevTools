package compress

type Compress interface {
	Compress(input string, output string) error
}

type CompressImpl struct {
	Quality int
}

func (obj *CompressImpl) Compress(input string, output string) error {
	return nil
}
