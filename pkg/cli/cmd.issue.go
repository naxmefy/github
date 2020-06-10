package cli

import cli "github.com/jawher/mow.cli"

func cmdIssue(cmd *cli.Cmd) {
	cmd.Command("create c", "create new issue", cmdIssueCreate)
	cmd.Command("list l", "list issues", cmdIssueList)
	cmd.Command("get g", "get issue info", cmdIssueGet)
	cmd.Command("set s", "update issue info", cmdIssueSet)
	cmd.Command("remove rm", "delete issue", cmdIssueRemove)

	cmd.Action = func() {
		cmd.PrintHelp()
	}
}

func cmdIssueCreate(cmd *cli.Cmd) {

}

func cmdIssueList(cmd *cli.Cmd) {

}

func cmdIssueGet(cmd *cli.Cmd) {

}

func cmdIssueSet(cmd *cli.Cmd) {

}

func cmdIssueRemove(cmd *cli.Cmd) {

}

