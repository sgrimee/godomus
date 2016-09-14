package cmd

import (
	"fmt"
	"os"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Get users",
	Long:  `Get users`,
	Run: func(cmd *cobra.Command, args []string) {
		validateSiteSet()
		sk := godomus.NewSiteKey(viper.GetInt("site"))
		users, err := domus.GetUsers(sk)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		if len(users) < 1 {
			fmt.Println("No users found, is the site correct?")
			os.Exit(-1)
		}
		output(outputFormat, users)

	},
}

func init() {
	getCmd.AddCommand(usersCmd)
}
