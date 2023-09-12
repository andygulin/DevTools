package convert

import (
	. "DevTools/tools/hash"
	"fmt"
	"github.com/spf13/cobra"
)

var obj Hash

func init() {
	obj = new(HashImpl)
}

var MD5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "Output MD5.",
	Long:  "Output the MD5 value based on the input text.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := obj.MD5(args[0])
		fmt.Println(str)
	},
}

var SHA1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "Output SHA1.",
	Long:  "Output the SHA1 value based on the input text.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := obj.SHA1(args[0])
		fmt.Println(str)
	},
}

var SHA256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "Output SHA256.",
	Long:  "Output the SHA256 value based on the input text.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := obj.SHA256(args[0])
		fmt.Println(str)
	},
}

var SHA512Cmd = &cobra.Command{
	Use:   "sha512",
	Short: "Output SHA512.",
	Long:  "Output the SHA512 value based on the input text.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := obj.SHA512(args[0])
		fmt.Println(str)
	},
}
