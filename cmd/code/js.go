package code

import (
	"log"
	fm "maxx/filemanager"

	"github.com/spf13/cobra"
)

var (
	js1 bool
	js2 bool
	js3 bool
	js4 bool
)

// jsCmd represents the js command
var jsCmd = &cobra.Command{
	Use:   "js",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// need quick reference for template files
		if len(args) != 0 && args[0] != "" {
			if js1 {
				fm.CopyFile("./templates/html/t1.html", args[0]+".html")
			} else if js2 {
				fm.CopyFile("./templates/html/t2.html", args[0]+".html")
			} else if js3 {
				fm.CopyFile("./templates/html/t3.html", args[0]+".html")
			} else if js4 {
				fm.CopyFile("./templates/html/t4.html", args[0]+".html")
			}
		} else {
			log.Println(" [ERROR] - Please name your file! `./maxx html --t# {name of file}`")
		}
	},
}

func init() {
	jsCmd.Flags().BoolVar(&js1, "t1", false, "Creates your t1.html file template")
	jsCmd.Flags().BoolVar(&js2, "t2", false, "Creates your t2.html file template")
	jsCmd.Flags().BoolVar(&js3, "t3", false, "Creates your t3.html file template")
	jsCmd.Flags().BoolVar(&js4, "t4", false, "Creates your t4.html file template")
	jsCmd.MarkFlagsMutuallyExclusive("t1", "t2", "t3", "t4")
	CodeCmd.AddCommand(jsCmd)
}
