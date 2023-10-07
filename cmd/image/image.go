package image

import (
	. "DevTools/tools/image/compress"
	. "DevTools/tools/image/convert"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var cp Compress
var cv Convert

func init() {
	cv = new(ConvertImpl)
}

var ImageCmd = &cobra.Command{
	Use:   "image",
	Short: "Image Compress/Convert.",
	Long:  "Image Compress/Convert.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var ImageCompressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Image Compress.",
	Long:  "Image Compress.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		quality, _ := strconv.Atoi(args[1])
		cp = &CompressImpl{Quality: quality}
		output, _ := cp.Compress(args[0])
		fmt.Println(output)
	},
}

var ImageConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Image Convert.",
	Long:  "Image Convert.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cv.Convert(args[0])
		fmt.Println(output)
	},
}
