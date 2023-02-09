package code

import (
	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var CodeCmd = &cobra.Command{
	Use:   "code",
	Short: "The 'code' pallette is used for setting up quick work environments",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
