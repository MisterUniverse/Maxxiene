/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package proc

import (
	"fmt"
	"maxx/pkg/procmgr"
	"os"

	"github.com/spf13/cobra"
)

// injectCmd represents the inject command
var injectCmd = &cobra.Command{
	Use:   "inject",
	Short: "Inject DLL into target application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inject called")
		if len(args) < 2 {
			fmt.Println("Error: You must provide exactly two parameters.")
			os.Exit(1) // Exit the program with an error code
		}

		target := args[0]
		payload := args[1]

		// Gets processid returns a unint32
		pID, err := procmgr.ProcessID(target)
		if err != nil {
			fmt.Println(err)
			return
		}

		// finds a process in the list by name and injects code
		err = procmgr.InjectDLL(pID, payload)
		if err != nil {
			fmt.Println("Error injecting DLL:", err)
		}
	},
}

func init() {
	ProcCmd.AddCommand(injectCmd)
}
