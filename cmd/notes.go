package cmd

import (
	"fmt"
	"maxx/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func createNewNote(path, s string) {
	if !exists(path) {
		f, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
	}
	pshCmd := fmt.Sprintf(
		"start notepad.exe ./notes/%v.md",
		s,
	)
	psh := utils.NewShell()

	stdOut, stdErr, err := psh.Execute(pshCmd)
	fmt.Printf("\n notes: \n StdOut : '%s' \n StdErr: '%s' \n Err: %s", strings.TrimSpace(stdOut), stdErr, err)

}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func printNoteArchive() {
	notes, _ := os.Open("./notes")

	files, err := notes.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	count := 0
	for _, v := range files {
		count++
		fmt.Printf("%v. %v\n", count, v.Name())
	}
}

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = os.Mkdir("./notes", 0755)

		if len(args) > 0 {
			createNewNote(fmt.Sprintf("./notes/%v.md", args[0]), args[0])
		} else {
			printNoteArchive()
		}
	},
}

func init() {
	rootCmd.AddCommand(notesCmd)
}
