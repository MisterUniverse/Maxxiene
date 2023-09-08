/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"maxx/pkg/filemgr"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a file or directory",
	Run:   runBackup,
}

func runBackup(cmd *cobra.Command, args []string) {
	config, _ := cmd.Flags().GetBool("config")
	if config {
		backupAllFromMap()
		return
	}

	if len(args) < 1 {
		fmt.Println("You must specify a file or directory to backup")
		return
	}
	backupSingle(args[0])
}

func backupAllFromMap() {
	paths := map[string]string{
		"todo":   viper.GetString("TODO_FILE_PATH"),
		"config": viper.GetString("CONFIG_DIR"),
		"data":   viper.GetString("DATA_DIR"),
		// more files or directories
	}

	for _, path := range paths {
		if err := backup(path); err != nil {
			fmt.Printf("Failed to backup %s: %s\n", path, err)
		}
	}
}

func backupSingle(path string) {
	if err := backup(path); err != nil {
		fmt.Printf("Failed to backup %s: %s\n", path, err)
	}
}

func backup(path string) error {
	timestamp := time.Now().Format("20060102-150405")
	backupFileName := filepath.Join(viper.GetString("BACKUPS"), fmt.Sprintf("%s-%s", filepath.Base(path), timestamp))

	var backupable filemgr.Backupable

	if filemgr.IsDir(path) {
		backupable = filemgr.NewDirBackup(path, fmt.Sprintf("%s.zip", backupFileName))
	} else {
		backupable = filemgr.NewFileBackup(path, backupFileName+".mxbkup")
	}

	if err := backupable.Backup(); err != nil {
		return err
	}
	fmt.Println("Backup successful:", backupFileName)
	return nil
}

func init() {
	backupCmd.Flags().BoolP("config", "c", false, "Backup every value from a predefined map")
	rootCmd.AddCommand(backupCmd)
	viper.SetConfigFile("config/.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}

}
