package cli

import cli "github.com/jawher/mow.cli"

func cmdRepo(cmd *cli.Cmd)  {
	cmd.Command("create c", "create new repository", cmdRepoCreate)
	cmd.Command("list l", "list repositories", cmdRepoList)
	cmd.Command("get g", "get repository info", cmdRepoGet)
	cmd.Command("set s", "update repository info", cmdRepoSet)
	cmd.Command("remove rm", "delete repository", cmdRepoRemove)

	cmd.Command("issue i", "several issue actions", cmdIssue)
	cmd.Command("pullrequest pr", "several pull request actions", cmdPullRequest)
	cmd.Command("label l", "several label actions", cmdLabel)
	cmd.Command("milestone ms", "several milestone actions", cmdMilestone)

	cmd.Action = func() {
		cmd.PrintHelp()
	}
}

func cmdRepoCreate(cmd *cli.Cmd) {

}

func cmdRepoList(cmd *cli.Cmd) {

}

func cmdRepoGet(cmd *cli.Cmd) {

}

func cmdRepoSet(cmd *cli.Cmd) {

}

func cmdRepoRemove(cmd *cli.Cmd) {

}
