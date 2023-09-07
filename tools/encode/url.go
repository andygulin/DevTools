package encode

import nu "net/url"

type Url interface {
	Encode(input string) (string, error)
	Decode(input string) (string, error)
}

type UrlImpl struct {
}

func (obj UrlImpl) Encode(input string) (string, error) {
	return nu.QueryEscape(input), nil
}

func (obj UrlImpl) Decode(input string) (string, error) {
	output, err := nu.QueryUnescape(input)
	if err != nil {
		return "", err
	}
	return output, nil
}
