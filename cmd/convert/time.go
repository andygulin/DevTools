package convert

import (
	. "DevTools/tools/convert/time"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var obj Time

func init() {
	obj = new(TimeImpl)
}

var TimeCmd = &cobra.Command{
	Use:   "time",
	Short: "Time Convert.",
	Long:  "Some development tools about time processing.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var TimeTimestampCmd = &cobra.Command{
	Use:   "timestamp",
	Short: "Current time.",
	Long:  "Print the current time in timestamp form: 1694170226.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := obj.Timestamp()
		fmt.Println(str)
	},
}

var TimeNowCmd = &cobra.Command{
	Use:   "now",
	Short: "Current time.",
	Long:  "The current time is printed as a string of 2020-01-01 22:22:22.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := obj.Now()
		fmt.Println(str)
	},
}

var TimeFormatTimeCmd = &cobra.Command{
	Use:   "format",
	Short: "Timestamp converts the time string",
	Long:  "Input a timestamp and output the converted time string in the format of 2023-01-01 22:22:22.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := strconv.ParseInt(args[0], 10, 64)
		str, _ := obj.FormatTime(input)
		fmt.Println(str)
	},
}
