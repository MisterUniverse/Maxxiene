package server

import (
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Spool up a server rigth quick.",
	Long: `The server command allows you to quicly spool up different types of servers.
For example:
- http
- tcp
- udp
- etc...`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}
