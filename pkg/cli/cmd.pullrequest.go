package cli

import cli "github.com/jawher/mow.cli"

func cmdPullRequest(cmd *cli.Cmd)  {
	cmd.Command("create c", "create new pull request", cmdPullRequestCreate)
	cmd.Command("list l", "list pull requests", cmdPullRequestList)
	cmd.Command("get g", "get pull request info", cmdPullRequestGet)
	cmd.Command("set s", "update pull request info", cmdPullRequestSet)
	cmd.Command("remove rm", "delete pull request", cmdPullRequestRemove)

	cmd.Action = func() {
		cmd.PrintHelp()
	}
}

func cmdPullRequestCreate(cmd *cli.Cmd) {

}

func cmdPullRequestList(cmd *cli.Cmd) {

}

func cmdPullRequestGet(cmd *cli.Cmd) {

}

func cmdPullRequestSet(cmd *cli.Cmd) {

}

func cmdPullRequestRemove(cmd *cli.Cmd) {

}
