package cli

import "github.com/jawher/mow.cli"

func App() *cli.Cli {
	app := cli.App("github", "github cli tool")

	return app
}
