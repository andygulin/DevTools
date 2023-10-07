package convert

import (
	. "DevTools/tools/convert/number"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var nn Number

func init() {
	nn = new(NumberImpl)
}

var NumberCmd = &cobra.Command{
	Use:   "number",
	Short: "base-conversion.",
	Long:  "Input binary, octal, decimal, hexadecimal, and output the converted value.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var Number2Cmd = &cobra.Command{
	Use:   "2",
	Short: "Binary",
	Long:  "Input binary, output converted value.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		ret8, _ := nn.BinaryToOctal(input)
		ret10, _ := nn.BinaryToDecimal(input)
		ret16, _ := nn.BinaryToHexadecimal(input)
		fmt.Printf("Binary : %s\n", input)
		fmt.Printf("Octal : %s\n", ret8)
		fmt.Printf("Decimal : %d\n", ret10)
		fmt.Printf("Hexadecimal : %s\n", ret16)
	},
}

var Number8Cmd = &cobra.Command{
	Use:   "8",
	Short: "Octal",
	Long:  "Input octal, output converted value.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		ret2, _ := nn.OctalToBinary(input)
		ret10, _ := nn.OctalToDecimal(input)
		ret16, _ := nn.OctalToHexadecimal(input)
		fmt.Printf("Octal : %s\n", input)
		fmt.Printf("Binary : %s\n", ret2)
		fmt.Printf("Decimal : %d\n", ret10)
		fmt.Printf("Hexadecimal : %s\n", ret16)
	},
}

var Number10Cmd = &cobra.Command{
	Use:   "10",
	Short: "Decimal",
	Long:  "Input decimal, output converted value.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := strconv.ParseInt(args[0], 10, 64)
		ret2, _ := nn.DecimalToBinary(input)
		ret8, _ := nn.DecimalToOctal(input)
		ret16, _ := nn.DecimalToHexadecimal(input)
		fmt.Printf("Decimal : %d\n", input)
		fmt.Printf("Binary : %s\n", ret2)
		fmt.Printf("Octal : %s\n", ret8)
		fmt.Printf("Hexadecimal : %s\n", ret16)
	},
}

var Number16Cmd = &cobra.Command{
	Use:   "16",
	Short: "Hexadecimal",
	Long:  "Input hexadecimal, output converted value.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		ret2, _ := nn.HexadecimalToBinary(input)
		ret8, _ := nn.HexadecimalToOctal(input)
		ret10, _ := nn.HexadecimalToDecimal(input)
		fmt.Printf("Hexadecimal : %s\n", input)
		fmt.Printf("Binary : %s\n", ret2)
		fmt.Printf("Octal : %s\n", ret8)
		fmt.Printf("Decimal : %d\n", ret10)
	},
}
