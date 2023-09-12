package encode

import (
	. "DevTools/tools/encode/url"
	"fmt"
	"github.com/spf13/cobra"
)

var obj Url

func init() {
	obj = new(UrlImpl)
}

var UrlCmd = &cobra.Command{
	Use:   "url",
	Short: "",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var UrlEncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := obj.Encode(args[0])
		fmt.Println(str)
	},
}

var UrlDecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := obj.Decode(args[0])
		fmt.Println(str)
	},
}
