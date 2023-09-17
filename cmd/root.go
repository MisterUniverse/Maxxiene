/*
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"maxx/cmd/convert"
	"maxx/cmd/crypto"
	"maxx/cmd/firewall"
	"maxx/cmd/notes"
	"maxx/cmd/proc"
	"maxx/cmd/sites"
	"maxx/cmd/task"
	mdb "maxx/pkg/db"
	"maxx/pkg/filemgr"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var maxxWorkingDir string // Default working dir for maxxiene

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "maxx",
	Short: "A productivity tool",
	Long:  `Maxxiene is a personalized cli to help with day to day task.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("calling from RUN")
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommand() {
	rootCmd.AddCommand(crypto.CryptoCmd)
	rootCmd.AddCommand(task.TaskCmd)
	rootCmd.AddCommand(sites.SitesCmd)
	rootCmd.AddCommand(convert.ConvertCmd)
	rootCmd.AddCommand(proc.ProcCmd)
	rootCmd.AddCommand(notes.NotesCmd)
	rootCmd.AddCommand(firewall.FirewallCmd)
}

func init() {
	initAppEnvironment()

	cobra.OnInitialize(initConfig)

	addSubCommand()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/Maxxiene/config/maxx.toml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Initialize the application environment.
func initAppEnvironment() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	maxxWorkingDir = fmt.Sprintf("%s\\Maxxiene", home)

	if !filemgr.IsDir(maxxWorkingDir) {
		fmt.Println("Maxxiene directory does not exist. Creating...")
		createWorkEnv()
	}
}

// initConfig reads in the configuration file and ENV variables if set.
func initConfig() {
	// // Use config file from the flag, if provided.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		setDefaultViperConfig()
	}

	// Automatically read in environment variables that match the config.
	viper.AutomaticEnv()

	// Attempt to read the configuration file.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintf(os.Stderr, "Using Maxxiene's default config file: %s\n", viper.ConfigFileUsed())
	}
}

// Set default Viper configuration options.
func setDefaultViperConfig() {
	viper.AddConfigPath(fmt.Sprintf("%s\\config", maxxWorkingDir))
	viper.SetConfigType("toml")
	viper.SetConfigName("maxx")
}

// createWorkingDirectories creates necessary directories for the application.
func createWorkingDirectories() {
	directories := []string{
		maxxWorkingDir,
		fmt.Sprintf("%s\\backups", maxxWorkingDir),
		fmt.Sprintf("%s\\config", maxxWorkingDir),
		fmt.Sprintf("%s\\data", maxxWorkingDir),
	}

	for _, dir := range directories {
		if err := filemgr.CreateDirectory(dir); err != nil {
			fmt.Println("Error creating directory:", err)
		}
	}
}

// createWorkEnv sets up the initial application environment.
func createWorkEnv() {
	fmt.Println("Initializing Maxxiene configuration and setup...")
	createWorkingDirectories()
	createInitialConfig()
	initDatabase()
	createRequiredFiles()
	fmt.Println("Configuration setup has been completed.")
}

// createInitialConfig generates the initial TOML configuration file.
func createInitialConfig() {
	key, err := GenerateRandomAESKey()
	if err != nil {
		fmt.Println("Error generating AES key:", err)
		return
	}

	keyHex := hex.EncodeToString(key)

	// Initialize settings.
	settings := map[string]map[string]string{
		"commands": {"listDataType": "note"},
		"log":      {"level": "info"},
		"crypto":   {"aes_key": keyHex},
		"paths": {
			"CONFIG":    fmt.Sprintf("%s\\config\\maxx.toml", maxxWorkingDir),
			"DATA_DIR":  fmt.Sprintf("%s\\data", maxxWorkingDir),
			"BOOKMARKS": fmt.Sprintf("%s\\data\\bookmarks.json", maxxWorkingDir),
			"BACKUPS":   fmt.Sprintf("%s\\backups", maxxWorkingDir),
			"DATABASE":  fmt.Sprintf("%s\\data\\maxxdb.db", maxxWorkingDir),
		},
	}

	err = filemgr.CreateTOMLFile(settings, settings["paths"]["CONFIG"])
	cobra.CheckErr(err)
}

// initDatabase initializes the database for the application.
func initDatabase() {
	fmt.Println("Initializing database...")
	mdb.MaxxDB = mdb.MaxxDataBase{}
	mdb.MaxxDB.Storage = mdb.NewDataStorage(maxxWorkingDir + "\\data\\maxxdb.db")
	mdb.MaxxDB.Storage.InitializeTables()
	fmt.Println("Database initialized!")
}

// createRequiredFiles creates any files required for the application to run.
func createRequiredFiles() {
	files := []string{
		fmt.Sprintf("%s\\data\\bookmarks.json", maxxWorkingDir),
	}

	for _, file := range files {
		if err := filemgr.CreateFile(file); err != nil {
			fmt.Println("Error creating file:", err)
		}
	}
}

// GenerateRandomAESKey generates a random 32-byte value for AES encryption.
func GenerateRandomAESKey() ([]byte, error) {
	key := make([]byte, 32) // 256 bits
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
