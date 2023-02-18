package code

import (
	fm "maxx/filemanager"
	"os"

	"github.com/spf13/cobra"
)

var (
	js1 bool
	js2 bool
	js3 bool
	js4 bool
)

func getJSTemplateFlag() string {
	if js1 {
		return "js_t1"
	} else if js2 {
		return "js_t2"
	} else if js3 {
		return "js_t3"
	} else if js4 {
		return "js_t4"
	} else {
		return ""
	}
}

// jsCmd represents the js command
var jsCmd = &cobra.Command{
	Use:   "js",
	Short: "This command will generate some boiler plate js",
	Long: `** js cmd description **
==============================================================================================================

Creates boiler plate js code based on the "t1.js-t4.js" files located in maxx's template directory.
Maxxiene can handle a total of 4 template .js files (this will be better in the future).
Each template file has to have this SPECIFIC naming convention:

- t1.js
- t2.js
- t3.js
- t4.js
	
(This will also change and improve in the future).`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && args[0] != "" {
			filename := args[0] + ".js"
			if flag := getJSTemplateFlag(); flag != "" {
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
	jsCmd.Flags().BoolVar(&js1, "t1", false, "Creates your t1.js file template")
	jsCmd.Flags().BoolVar(&js2, "t2", false, "Creates your t2.js file template")
	jsCmd.Flags().BoolVar(&js3, "t3", false, "Creates your t3.js file template")
	jsCmd.Flags().BoolVar(&js4, "t4", false, "Creates your t4.js file template")
	jsCmd.MarkFlagsMutuallyExclusive("t1", "t2", "t3", "t4")
	CodeCmd.AddCommand(jsCmd)
}
