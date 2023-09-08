package procmgr

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/process"
	"golang.org/x/sys/windows"
)

// GetPIDByName returns the process IDs based on the process name.
func GetPIDByName(name string) ([]int32, error) {
	var pids []int32
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	for _, p := range processes {
		pName, err := p.Name()
		if err != nil {
			continue
		}
		if pName == name {
			pids = append(pids, p.Pid)
		}
	}
	return pids, nil
}

// GetProcessInfoByPID returns CPU and Memory information for a given PID.
func GetProcessInfoByPID(pid int32) (float64, float32, error) {
	p, err := process.NewProcess(pid)
	if err != nil {
		return 0, 0, err
	}

	cpuPercent, err := p.CPUPercent()
	if err != nil {
		return 0, 0, err
	}

	memoryPercent, err := p.MemoryPercent()
	if err != nil {
		return 0, 0, err
	}

	return cpuPercent, memoryPercent, nil
}

// KillProcess kills a process by its PID
func KillProcess(pid int) error {
	handle, err := windows.OpenProcess(windows.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		fmt.Printf("Error opening process handle: %v", err)
		return err
	}
	defer windows.CloseHandle(handle)

	exitCode := uint32(1)
	err = windows.TerminateProcess(handle, exitCode)
	if err != nil {
		fmt.Printf("Error terminating process: %v", err)
		return err
	}

	fmt.Printf("Successfully terminated process with PID %d", pid)
	return nil
}

// WatchProcess monitors CPU and Memory usage for a given PID.
func WatchProcess(pid int32, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cpuPercent, memoryPercent, err := GetProcessInfoByPID(pid)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Printf("CPU: %.2f%%, Memory: %.2f%%\n", cpuPercent, memoryPercent)
		}
	}
}
