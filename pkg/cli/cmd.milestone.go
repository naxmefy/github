package cli

import cli "github.com/jawher/mow.cli"

func cmdMilestone(cmd *cli.Cmd) {
	cmd.Command("create c", "create new milestone", cmdMilestoneCreate)
	cmd.Command("list l", "list milestones", cmdMilestoneList)
	cmd.Command("get g", "get milestone info", cmdMilestoneGet)
	cmd.Command("set s", "update milestone info", cmdMilestoneSet)
	cmd.Command("remove rm", "delete milestone", cmdMilestoneRemove)

	cmd.Action = func() {
		cmd.PrintHelp()
	}
}

func cmdMilestoneCreate(cmd *cli.Cmd) {

}

func cmdMilestoneList(cmd *cli.Cmd) {

}

func cmdMilestoneGet(cmd *cli.Cmd) {

}

func cmdMilestoneSet(cmd *cli.Cmd) {

}

func cmdMilestoneRemove(cmd *cli.Cmd) {

}

