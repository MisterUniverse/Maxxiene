package code

import (
	fm "maxx/filemanager"
	"os"

	"github.com/spf13/cobra"
)

var (
	go1 bool
	go2 bool
	go3 bool
	go4 bool
)

func getGOTemplateFlag() string {
	if go1 {
		return "golang_t1"
	} else if go2 {
		return "golang_t2"
	} else if go3 {
		return "golang_t3"
	} else if go4 {
		return "golang_t4"
	} else {
		return ""
	}
}

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "This command will generate some boiler plate go code",
	Long: `** go cmd description **
==============================================================================================================

Creates boiler plate go code based on the "t1.go-t4.go" files located in maxx's template directory.
Maxxiene can handle a total of 4 template .go files (this will be better in the future).
Each template file has to have this SPECIFIC naming convention:

- t1.go
- t2.go
- t3.go
- t4.go
	
(This will also change and improve in the future).`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && args[0] != "" {
			filename := args[0] + ".go"
			if flag := getGOTemplateFlag(); flag != "" {
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
	goCmd.Flags().BoolVar(&go1, "t1", false, "Creates your t1.html file template")
	goCmd.Flags().BoolVar(&go2, "t2", false, "Creates your t2.html file template")
	goCmd.Flags().BoolVar(&go3, "t3", false, "Creates your t3.html file template")
	goCmd.Flags().BoolVar(&go4, "t4", false, "Creates your t4.html file template")
	goCmd.MarkFlagsMutuallyExclusive("t1", "t2", "t3", "t4")
	CodeCmd.AddCommand(goCmd)
}
