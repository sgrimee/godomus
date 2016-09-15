// Generate a bash completion script
// Not working for me at the moment...

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
