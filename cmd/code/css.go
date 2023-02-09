package code

import (
	"log"
	fm "maxx/filemanager"

	"github.com/spf13/cobra"
)

var (
	css1 bool
	css2 bool
	css3 bool
	css4 bool
)

var cssCmd = &cobra.Command{
	Use:   "css",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// need quick reference for template files
		if len(args) != 0 && args[0] != "" {
			if css1 {
				fm.CopyFile("./templates/html/t1.html", args[0]+".html")
			} else if css2 {
				fm.CopyFile("./templates/html/t2.html", args[0]+".html")
			} else if css3 {
				fm.CopyFile("./templates/html/t3.html", args[0]+".html")
			} else if css4 {
				fm.CopyFile("./templates/html/t4.html", args[0]+".html")
			}
		} else {
			log.Println(" [ERROR] - Please name your file! `./maxx html --t# {name of file}`")
		}
	},
}

func init() {
	cssCmd.Flags().BoolVar(&css1, "t1", false, "Creates your t1.html file template")
	cssCmd.Flags().BoolVar(&css2, "t2", false, "Creates your t2.html file template")
	cssCmd.Flags().BoolVar(&css3, "t3", false, "Creates your t3.html file template")
	cssCmd.Flags().BoolVar(&css4, "t4", false, "Creates your t4.html file template")
	cssCmd.MarkFlagsMutuallyExclusive("t1", "t2", "t3", "t4")
	CodeCmd.AddCommand(cssCmd)
}
