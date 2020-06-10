package cli

import cli "github.com/jawher/mow.cli"

func cmdNotification(cmd *cli.Cmd) {
	cmd.Command("list l", "list notifications", cmdNotificationList)
	cmd.Command("get g", "get notification info", cmdNotificationGet)
	cmd.Command("set s", "update notification info", cmdNotificationSet)
	cmd.Command("remove rm", "delete notification", cmdNotificationRemove)

	cmd.Action = func() {
		cmd.PrintHelp()
	}
}

func cmdNotificationList(cmd *cli.Cmd) {

}

func cmdNotificationGet(cmd *cli.Cmd) {

}

func cmdNotificationSet(cmd *cli.Cmd) {

}

func cmdNotificationRemove(cmd *cli.Cmd) {

}
