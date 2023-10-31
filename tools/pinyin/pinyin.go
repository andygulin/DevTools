package pinyin

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
)

type Pinyin interface {
	Convert(input string) (string, error)
}

type PinyinImpl struct {
}

func (obj *PinyinImpl) Convert(input string) (string, error) {
	args := pinyin.NewArgs()
	args.Style = pinyin.Normal
	output := pinyin.Pinyin(input, args)
	var rets []string
	for _, str := range output {
		rets = append(rets, str[0])
	}
	return strings.Join(rets, ""), nil
}
