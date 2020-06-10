package cli

import cli "github.com/jawher/mow.cli"

func cmdLabel(cmd *cli.Cmd) {
	cmd.Command("create c", "create new label", cmdLabelCreate)
	cmd.Command("list l", "list labels", cmdLabelList)
	cmd.Command("get g", "get label info", cmdLabelGet)
	cmd.Command("set s", "update label info", cmdLabelSet)
	cmd.Command("remove rm", "delete label", cmdLabelRemove)

	cmd.Action = func() {
		cmd.PrintHelp()
	}
}

func cmdLabelCreate(cmd *cli.Cmd) {

}

func cmdLabelList(cmd *cli.Cmd) {

}

func cmdLabelGet(cmd *cli.Cmd) {

}

func cmdLabelSet(cmd *cli.Cmd) {

}

func cmdLabelRemove(cmd *cli.Cmd) {

}

