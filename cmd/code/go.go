package code

import (
	"log"
	fm "maxx/filemanager"

	"github.com/spf13/cobra"
)

var (
	go1 bool
	go2 bool
	go3 bool
	go4 bool
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Creates some boiler plate golang code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// need quick reference for template files
		if len(args) != 0 && args[0] != "" {
			if go1 {
				fm.CopyFile("./templates/html/t1.html", args[0]+".html")
			} else if go2 {
				fm.CopyFile("./templates/html/t2.html", args[0]+".html")
			} else if go3 {
				fm.CopyFile("./templates/html/t3.html", args[0]+".html")
			} else if go4 {
				fm.CopyFile("./templates/html/t4.html", args[0]+".html")
			}
		} else {
			log.Println(" [ERROR] - Please name your file! `./maxx html --t# {name of file}`")
		}
	},
}

func init() {
	goCmd.Flags().BoolVar(&go1, "t1", false, "Creates your t1.html file template")
	goCmd.Flags().BoolVar(&go2, "t2", false, "Creates your t2.html file template")
	goCmd.Flags().BoolVar(&go3, "t3", false, "Creates your t3.html file template")
	goCmd.Flags().BoolVar(&go4, "t4", false, "Creates your t4.html file template")
	goCmd.MarkFlagsMutuallyExclusive("t1", "t2", "t3", "t4")
	CodeCmd.AddCommand(goCmd)
}
