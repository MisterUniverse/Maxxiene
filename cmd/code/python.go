package code

import (
	fm "maxx/filemanager"
	"os"

	"github.com/spf13/cobra"
)

var (
	py1 bool
	py2 bool
	py3 bool
	py4 bool
)

func getPYTemplateFlag() string {
	if py1 {
		return "python_t1"
	} else if py2 {
		return "python_t2"
	} else if py3 {
		return "python_t3"
	} else if py4 {
		return "python_t4"
	} else {
		return ""
	}
}

// pythonCmd represents the python command
var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "This command will generate some boiler plate python code",
	Long: `** python cmd description **
==============================================================================================================

Creates boiler plate python code based on the "t1.py-t4.py" files located in maxx's template directory.
Maxxiene can handle a total of 4 template .py files (this will be better in the future).
Each template file has to have this SPECIFIC naming convention:

- t1.py
- t2.py
- t3.py
- t4.py
	
(This will also change and improve in the future).`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && args[0] != "" {
			filename := args[0] + ".py"
			if flag := getPYTemplateFlag(); flag != "" {
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
	pythonCmd.Flags().BoolVar(&py1, "t1", false, "Creates your t1.html file template")
	pythonCmd.Flags().BoolVar(&py2, "t2", false, "Creates your t2.html file template")
	pythonCmd.Flags().BoolVar(&py3, "t3", false, "Creates your t3.html file template")
	pythonCmd.Flags().BoolVar(&py4, "t4", false, "Creates your t4.html file template")
	pythonCmd.MarkFlagsMutuallyExclusive("t1", "t2", "t3", "t4")
	CodeCmd.AddCommand(pythonCmd)
}
