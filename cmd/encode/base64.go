package encode

import (
	. "DevTools/tools/encode/base64"
	"fmt"
	"github.com/spf13/cobra"
)

var bb Base64

func init() {
	bb = new(Base64Impl)
}

var Base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "base64 transcoding.",
	Long:  "Input string or image, output base64 encoding.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var Base64TextCmd = &cobra.Command{
	Use:   "text",
	Short: "String transcoding.",
	Long:  "Input string, output base64 encoding.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := bb.EncodeText(TypeText(args[0]))
		fmt.Println(str)
	},
}

var Base64ImageCmd = &cobra.Command{
	Use:   "image",
	Short: "Image transcoding.",
	Long:  "Input image, output base64 encoding.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := bb.EncodeImage(TypeImageFile(args[0]))
		fmt.Println(str)
	},
}
