package convert

import (
	. "DevTools/tools/qrcode"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var obj QrCode

var QrCodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "Generate QrCode.",
	Long:  "Generate QrCode.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := strconv.Atoi(args[1])
		obj = &QrCodeImpl{Size: size}
		str, _ := obj.Create(args[0])
		fmt.Println(str)
	},
}
