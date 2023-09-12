package convert

import (
	. "DevTools/tools/uuid"
	"fmt"
	"github.com/spf13/cobra"
)

var obj UUID

func init() {
	obj = new(GoogleUUID)
}

var UUIDCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Output UUID.",
	Long:  "Use google's library to output UUID.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := obj.GenerateUUID()
		fmt.Println(str)
	},
}
