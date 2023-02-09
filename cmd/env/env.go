package env

import (
	"github.com/spf13/cobra"
)

var EnvCmd = &cobra.Command{
	Use:   "env",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
