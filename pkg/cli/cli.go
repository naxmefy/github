package cli

import (
	"github.com/jawher/mow.cli"
	"github.com/naxmefy/github/pkg"
)

func App() *cli.Cli {
	app := cli.App("github", "github cli tool")
	app.Version("v version", pkg.VERSION)

	app.Command("login", "login into github", cmdLogin)
	app.Command("notification n", "several notification actions", cmdNotification)
	app.Command("repo r", "several repository actions", cmdRepo)
	app.Command("whoami", "shows current login information", cmdWhoAmI)

	// TODO
	app.Command("repl", "[unfinished] interactive shell", cmdRepl)
	app.Command("ui", "[unfinished] ui dashboard", cmdUi)

	app.Action = func() {
		app.PrintHelp()
	}

	return app
}
