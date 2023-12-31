package cmd

import (
	. "DevTools/cmd/convert"
	. "DevTools/cmd/encode"
	. "DevTools/cmd/format"
	. "DevTools/cmd/hash"
	. "DevTools/cmd/image"
	. "DevTools/cmd/passwd"
	. "DevTools/cmd/pinyin"
	. "DevTools/cmd/qrcode"
	. "DevTools/cmd/uuid"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "DevTools",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	TimeCmd.AddCommand(TimeTimestampCmd, TimeNowCmd, TimeFormatTimeCmd)

	rootCmd.AddCommand(TimeCmd)
	rootCmd.AddCommand(UUIDCmd)
	rootCmd.AddCommand(MD5Cmd, SHA1Cmd, SHA256Cmd, SHA512Cmd)

	UrlCmd.AddCommand(UrlEncodeCmd, UrlDecodeCmd)
	rootCmd.AddCommand(UrlCmd)

	Base64Cmd.AddCommand(Base64TextCmd, Base64ImageCmd)
	rootCmd.AddCommand(Base64Cmd)

	JwtCmd.AddCommand(JwtCreateCmd, JwtParseCmd)
	rootCmd.AddCommand(JwtCmd)

	NumberCmd.AddCommand(Number2Cmd, Number8Cmd, Number10Cmd, Number16Cmd)
	rootCmd.AddCommand(NumberCmd)

	ImageCmd.AddCommand(ImageCompressCmd, ImageConvertCmd)
	rootCmd.AddCommand(ImageCmd)

	rootCmd.AddCommand(Json2YamlCmd)

	rootCmd.AddCommand(FormatCmd)
	rootCmd.AddCommand(FormatFileCmd)

	rootCmd.AddCommand(PasswordCmd)

	rootCmd.AddCommand(PinyinCmd)

	rootCmd.AddCommand(QrCodeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
