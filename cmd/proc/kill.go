/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package proc

import (
	"fmt"
	"maxx/pkg/procmgr"

	"github.com/spf13/cobra"
)

// killCmd represents the kill command
var killCmd = &cobra.Command{
	Use:   "kill [process name]",
	Short: "Kill a process by its name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You must specify a process name.")
			return
		}

		processName := args[0]
		pids, err := procmgr.GetPIDByName(processName)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		if len(pids) == 0 {
			fmt.Printf("No processes found with the name: %s\n", processName)
			return
		}

		// Assume we kill the first process if multiple are found
		pid := pids[0]
		fmt.Printf("Killing process %s with PID %d\n", processName, pid)

		if err := procmgr.KillProcess(int(pid)); err != nil {
			fmt.Printf("Failed to kill process: %s\n", err)
		} else {
			fmt.Printf("Successfully killed process %s with PID %d\n", processName, pid)
		}
	},
}

func init() {
	ProcCmd.AddCommand(killCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// killCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// killCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
