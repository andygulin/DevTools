package encode

import (
	. "DevTools/tools/encode/jwt"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var jj Jwt

func init() {
	jwtKey := "u*29HJ1320"
	expireMinute := 30
	jj = &JwtImpl{
		JwtKey:       jwtKey,
		ExpireMinute: expireMinute,
	}
}

var JwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Jwt Create/Parse.",
	Long:  "Jwt Create/Parse.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var JwtCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Jwt Create.",
	Long:  "Jwt Create.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		userId, _ := strconv.Atoi(args[0])
		str, _ := jj.Create(userId)
		fmt.Println(str)
	},
}

var JwtParseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Jwt Parse.",
	Long:  "Jwt Parse.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		userId, err := jj.Parse(args[0])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(userId)
		}
	},
}
