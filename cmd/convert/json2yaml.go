package convert

import (
	. "DevTools/tools/convert/json2yaml"
	"fmt"
	"github.com/spf13/cobra"
)

var jy Json2Yaml

func init() {
	jy = new(Json2YamlImpl)
}

var Json2YamlCmd = &cobra.Command{
	Use:   "json2yaml",
	Short: "Json file to Yaml file.",
	Long:  "Json file to Yaml file.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := jy.Convert(args[0])
		fmt.Println(output)
	},
}
