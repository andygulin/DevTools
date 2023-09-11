package cmd

import (
	. "DevTools/cmd/convert"
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
	TimeCmd.AddCommand(TimeTimestampCmd)
	TimeCmd.AddCommand(TimeNowCmd)
	TimeCmd.AddCommand(TimeFormatTimeCmd)

	rootCmd.AddCommand(TimeCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
