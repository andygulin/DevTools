package convert

import (
	. "DevTools/tools/pinyin"
	"fmt"
	"github.com/spf13/cobra"
)

var obj Pinyin

func init() {
	obj = new(PinyinImpl)
}

var PinyinCmd = &cobra.Command{
	Use:   "pinyin",
	Short: "Chinese characters to pinyin.",
	Long:  "Chinese characters to pinyin.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := obj.Convert(args[0])
		fmt.Println(str)
	},
}
