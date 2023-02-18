package code

import (
	fm "maxx/filemanager"
	"os"

	"github.com/spf13/cobra"
)

var (
	css1 bool
	css2 bool
	css3 bool
	css4 bool
)

func getCSSTemplateFlag() string {
	if css1 {
		return "css_t1"
	} else if css2 {
		return "css_t2"
	} else if css3 {
		return "css_t3"
	} else if css4 {
		return "css_t4"
	} else {
		return ""
	}
}

var cssCmd = &cobra.Command{
	Use:   "css",
	Short: "This command will generate some boiler plate css",
	Long: `** css cmd description**
==============================================================================================================

Creates boiler plate css based on the "t1.css-t4.css" files located in maxx's template directory.
Maxxiene can handle a total of 4 template .css files (this will be better in the future).
Each template file has to have this SPECIFIC naming convention:

- t1.css
- t2.css
- t3.css
- t4.css
	
(This will also change and improve in the future).`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && args[0] != "" {
			filename := args[0] + ".css"
			if flag := getCSSTemplateFlag(); flag != "" {
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
	cssCmd.Flags().BoolVar(&css1, "t1", false, "Creates your t1.html file template")
	cssCmd.Flags().BoolVar(&css2, "t2", false, "Creates your t2.html file template")
	cssCmd.Flags().BoolVar(&css3, "t3", false, "Creates your t3.html file template")
	cssCmd.Flags().BoolVar(&css4, "t4", false, "Creates your t4.html file template")
	cssCmd.MarkFlagsMutuallyExclusive("t1", "t2", "t3", "t4")
	CodeCmd.AddCommand(cssCmd)
}
