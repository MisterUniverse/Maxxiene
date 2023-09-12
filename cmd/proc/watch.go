/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package proc

import (
	"fmt"
	"maxx/pkg/procmgr"
	"time"

	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch [process name]",
	Short: "Watch a process by its name",
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

		// Assume we watch the first process if multiple are found
		pid := pids[0]
		fmt.Printf("Watching process %s with PID %d\n", processName, pid)

		interval := 2 * time.Second // Monitoring interval
		procmgr.WatchProcess(pid, interval)
	},
}

func init() {
	ProcCmd.AddCommand(watchCmd)
}
