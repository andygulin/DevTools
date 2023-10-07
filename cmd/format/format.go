package format

import (
	. "DevTools/tools/format"
	. "DevTools/tools/format/json"
	. "DevTools/tools/format/sql"
	. "DevTools/tools/format/xml"
	"fmt"
	"github.com/spf13/cobra"
)

var xf Format
var sf Format
var js Format

func init() {
	xf = new(XmlFormat)
	sf = new(SQLFormat)
	js = new(JsonFormat)
}

var FormatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format Xml, Json, and SQL text.",
	Long:  "Format Xml, Json, and SQL text.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tp := args[0]
		input := args[1]
		var output string
		if tp == "xml" {
			output, _ = xf.Format(input)
		}
		if tp == "json" {
			output, _ = js.Format(input)
		}
		if tp == "sql" {
			output, _ = sf.Format(input)
		}

		fmt.Println(output)
	},
}

var FormatFileCmd = &cobra.Command{
	Use:   "format_file",
	Short: "Format Xml, Json, and SQL files.",
	Long:  "Format Xml, Json, and SQL files.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tp := args[0]
		input := args[1]
		var output string
		if tp == "xml" {
			output, _ = xf.FormatFile(input)
		}
		if tp == "json" {
			output, _ = js.FormatFile(input)
		}
		if tp == "sql" {
			output, _ = sf.FormatFile(input)
		}

		fmt.Println(output)
	},
}
