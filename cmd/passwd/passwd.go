package convert

import (
	. "DevTools/tools/passwd"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var obj Password

var PasswordCmd = &cobra.Command{
	Use:   "passwd",
	Short: "Generate random passwords.",
	Long:  "Generate random passwords.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := strconv.Atoi(args[0])
		obj = &PasswordImpl{Length: length}
		str, _ := obj.RandomPassword()
		fmt.Println(str)
	},
}
