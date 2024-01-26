/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

// unblockCmd represents the unblock command
var unblockCmd = &cobra.Command{
	Use:   "unblock [path]",
	Short: "Unblock a file or directory",
	Long:  `Unblock a file or directory and all its sub-folders and files.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := unblockPath(args[0]); err != nil {
			fmt.Println("Error unblocking:", err)
			os.Exit(1)
		}
		fmt.Println("Successfully unblocked:", args[0])
	},
}

func unblockPath(path string) error {
	var files []string
	const batchSize = 100 // Adjust batch size here

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 7) // Control concurrency level here

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		files = append(files, path)
		if len(files) >= batchSize {
			wg.Add(1)
			semaphore <- struct{}{}
			go unblockFiles(files, &wg, semaphore)
			files = nil // reset the file list
		}

		return nil
	})

	// Process remaining files
	if len(files) > 0 {
		wg.Add(1)
		semaphore <- struct{}{}
		go unblockFiles(files, &wg, semaphore)
	}

	wg.Wait()
	return err
}

func unblockFiles(files []string, wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()
	defer func() { <-semaphore }()

	// Join file paths into a single string for the PowerShell command
	filePaths := strings.Join(files, "', '")
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Get-Item -Path '%s' | Unblock-File", filePaths))
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to unblock files:", err)
	}
}

func init() {
	rootCmd.AddCommand(unblockCmd)
}
