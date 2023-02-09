package code

import (
	"log"
	fm "maxx/filemanager"

	"github.com/spf13/cobra"
)

var (
	html1 bool
	html2 bool
	html3 bool
	html4 bool
)

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// need quick reference for template files
		if len(args) != 0 && args[0] != "" {
			if html1 {
				fm.CopyFile("./templates/html/t1.html", args[0]+".html")
			} else if html2 {
				fm.CopyFile("./templates/html/t2.html", args[0]+".html")
			} else if html3 {
				fm.CopyFile("./templates/html/t3.html", args[0]+".html")
			} else if html4 {
				fm.CopyFile("./templates/html/t4.html", args[0]+".html")
			}
		} else {
			log.Println(" [ERROR] - Please name your file! `./maxx html --t# {name of file}`")
		}

	},
}

func init() {
	htmlCmd.Flags().BoolVar(&html1, "t1", false, "Creates your t1.html file template")
	htmlCmd.Flags().BoolVar(&html2, "t2", false, "Creates your t2.html file template")
	htmlCmd.Flags().BoolVar(&html3, "t3", false, "Creates your t3.html file template")
	htmlCmd.Flags().BoolVar(&html4, "t4", false, "Creates your t4.html file template")
	htmlCmd.MarkFlagsMutuallyExclusive("t1", "t2", "t3", "t4")
	CodeCmd.AddCommand(htmlCmd)
}
