package cli

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
)

func cmdRepl(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Println("not implemented yet")
	}
}
