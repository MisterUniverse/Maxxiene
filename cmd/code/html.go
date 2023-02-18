package code

import (
	fm "maxx/filemanager"
	"os"

	"github.com/spf13/cobra"
)

var (
	html1 bool
	html2 bool
	html3 bool
	html4 bool
)

func getHTMLTemplateFlag() string {
	if html1 {
		return "html_t1"
	} else if html2 {
		return "html_t2"
	} else if html3 {
		return "html_t3"
	} else if html4 {
		return "html_t4"
	} else {
		return ""
	}
}

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "This command will generate some boiler plate html",
	Long: `** html cmd description **
==============================================================================================================

Creates boiler plate html based on the "t1.html-t4.html" files located in maxx's template directory.
Maxxiene can handle a total of 4 template .html/htm files (this will be better in the future).
Each template file has to have this SPECIFIC naming convention:

- t1.html
- t2.html
- t3.html
- t4.html
	
(This will also change and improve in the future).`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && args[0] != "" {
			filename := args[0] + ".html"
			if flag := getHTMLTemplateFlag(); flag != "" {
				if tmpl, ok := templates.templates[flag]; ok {
					fm.CopyFile(tmpl, filename)
					cmd.Println("Generated file ", filename)
				} else {
					cmd.PrintErrf("Error: Invalid template flag: %s\n", flag)
					os.Exit(1)
				}
			} else {
				cmd.PrintErrln("Error: Please specify a template flag.")
				cmd.Usage()
				os.Exit(1)
			}
		} else {
			cmd.PrintErrln("Error: Please specify a filename.")
			cmd.Usage()
			os.Exit(1)
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
