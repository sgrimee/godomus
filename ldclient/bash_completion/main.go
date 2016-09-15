// Generate a bash completion script

package main

import (
	"log"
	"os"

	"github.com/sgrimee/godomus/ldclient/cmd"
)

func main() {
	file, err := os.OpenFile(
		"bash-completion.sh",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	cmd.RootCmd.GenBashCompletion(file)
}
