package code

import (
	fm "maxx/filemanager"

	"github.com/spf13/cobra"
)

var (
	t1 bool
	t2 bool
)

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// need quick reference for template files
		if args[0] != "" {
			if t1 {
				fm.CopyFile("./templates/html/index.html", args[0]+".html")
			} else if t2 {
				fm.CopyFile("./templates/html/divs.html", args[0]+".html")
			}
		}

	},
}

func init() {
	htmlCmd.Flags().BoolVar(&t1, "t1", false, "Creates a basic html boilerplate")
	htmlCmd.Flags().BoolVar(&t2, "t2", false, "Creates four divs")
	htmlCmd.MarkFlagsMutuallyExclusive("t1", "t2")
	CodeCmd.AddCommand(htmlCmd)
}
