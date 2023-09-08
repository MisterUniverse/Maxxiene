/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package proc

import (
	"fmt"
	"maxx/pkg/procmgr"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sys/windows"
)

// MiniDump constants
const (
	MiniDumpNormal         = 0x00000000
	MiniDumpWithDataSegs   = 0x00000001
	MiniDumpWithFullMemory = 0x00000002
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump [process name]",
	Short: "Create a memory dump of a process by its name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You must specify a process name.")
			return
		}

		processName := args[0]
		pids, err := procmgr.GetPIDByName(processName) // Assume you have implemented GetPIDByName

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		if len(pids) == 0 {
			fmt.Printf("No processes found with the name: %s\n", processName)
			return
		}

		pid := pids[0]
		fmt.Printf("Creating memory dump for process %s with PID %d\n", processName, pid)

		err = createMemoryDump(int(pid), processName)
		if err != nil {
			fmt.Printf("Failed to create memory dump: %s\n", err)
		} else {
			fmt.Printf("Successfully created memory dump for process %s with PID %d\n", processName, pid)
		}
	},
}

func createMemoryDump(pid int, name string) error {
	dbghelp := windows.NewLazySystemDLL("Dbghelp.dll")
	miniDumpWriteDump := dbghelp.NewProc("MiniDumpWriteDump")

	handle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ, false, uint32(pid))
	if err != nil {
		return err
	}
	defer windows.CloseHandle(handle)

	dumpFilePtr, err := windows.UTF16PtrFromString(viper.GetString("DATA_DIR") + "/" + name + ".dmp")
	if err != nil {
		return err
	}

	file, err := windows.CreateFile(dumpFilePtr, windows.GENERIC_WRITE, 0, nil, windows.CREATE_ALWAYS, windows.FILE_ATTRIBUTE_NORMAL, 0)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(file)

	r1, _, err := miniDumpWriteDump.Call(
		uintptr(handle),
		uintptr(pid),
		uintptr(file),
		uintptr(MiniDumpWithDataSegs),
		0,
		0,
		0,
	)
	if r1 == 0 {
		return err
	}

	return nil
}

func init() {
	viper.SetConfigFile("config/.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}

	ProcCmd.AddCommand(dumpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
