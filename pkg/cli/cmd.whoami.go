package cli

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
)

func cmdWhoAmI(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Println("Not logged in. Use 'github login'")
	}
}
