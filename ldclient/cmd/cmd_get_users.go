package cmd

import (
	"log"

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
		sk := godomus.NewSiteKey(viper.GetInt("site"))
		users, err := domus.GetUsers(sk)
		if err != nil {
			log.Fatal(err)
		}
		if len(users) < 1 {
			log.Fatal("No users found, is the site correct?")
		}
		output(outputFormat, users)

	},
}

func init() {
	getCmd.AddCommand(usersCmd)
}
