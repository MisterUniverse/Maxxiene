/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"maxx/pkg/filemgr"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a file: or directory",
	Run:   runBackup,
}

func runBackup(cmd *cobra.Command, args []string) {
	dst := viper.GetString("paths.BACKUPS")

	auto, _ := cmd.Flags().GetBool("auto")
	if auto {
		if len(args) < 2 {
			fmt.Println("Must enter path/to/cellardoor/settings.json and path/to/cellardoor/main.py")
			return
		}
		filemgr.CellarDoor(args[0], args[1])
		return
	}

	config, _ := cmd.Flags().GetBool("config")
	if config {
		paths := map[string]string{
			"config": viper.GetString("paths.CONFIG_DIR"),
			"data":   viper.GetString("paths.DATA_DIR"),
		}
		filemgr.BackupAllFromMap(paths, dst)
		return
	}

	if len(args) < 1 {
		fmt.Println("You must specify a file or directory to backup")
		return
	}
	filemgr.BackupSingle(args[0], dst)
}

func init() {
	backupCmd.Flags().BoolP("config", "c", false, "Backup every value from a predefined map")
	backupCmd.Flags().BoolP("auto", "a", false, "Auto backup using CellarDoor")
	rootCmd.AddCommand(backupCmd)
}
