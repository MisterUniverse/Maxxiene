package filemgr

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// Configuration and App structs to hold the JSON configuration
type App struct {
	DESTINATION string   `json:"DESTINATION"`
	GO_TIME     string   `json:"GO_TIME"`
	MAX_BACKUPS int      `json:"MAX_BACKUPS"`
	SOURCE      []string `json:"SOURCE"`
}

type Configuration struct {
	APP App `json:"APP"`
}

// readConfig reads the configuration from a file and returns it
func readConfig(fileName string) (Configuration, error) {
	var config Configuration
	file, err := os.ReadFile(fileName)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(file, &config)
	return config, err
}

// writeConfig writes the configuration to a file
func writeConfig(config Configuration, fileName string) error {
	updatedConfig, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, updatedConfig, 0644)
	return err
}

func CellarDoor(config, pyscript string) {
	updateCellarDoorSettings(config)
	executeCommand("python", pyscript)
}

// executeCommand executes a shell command and returns its output
func executeCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func updateCellarDoorSettings(fileName string) {
	// Read the existing configuration
	config, err := readConfig(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// Display and edit configuration
	fmt.Println("Current settings:")
	fmt.Printf("DESTINATION: %s\n", config.APP.DESTINATION)
	fmt.Printf("GO_TIME: %s\n", config.APP.GO_TIME)
	fmt.Printf("MAX_BACKUPS: %d\n", config.APP.MAX_BACKUPS)
	fmt.Printf("SOURCE: %v\n", config.APP.SOURCE)

	var input string

	fmt.Print("Update DESTINATION [Leave blank to keep]: ")
	fmt.Scanln(&input)
	if input != "" {
		config.APP.DESTINATION = input
	}

	fmt.Print("Update GO_TIME [Leave blank to keep]: ")
	fmt.Scanln(&input)
	if input != "" {
		config.APP.GO_TIME = input
	}

	fmt.Print("Update MAX_BACKUPS [Leave blank to keep]: ")
	fmt.Scanln(&input)
	if input != "" {
		var newMaxBackups int
		_, err := fmt.Sscanf(input, "%d", &newMaxBackups)
		if err == nil {
			config.APP.MAX_BACKUPS = newMaxBackups
		} else {
			fmt.Println("Invalid input, keeping old value.")
		}
	}

	fmt.Print("Update SOURCE (comma-separated, without spaces) [Leave blank to keep]: ")
	fmt.Scanln(&input)
	if input != "" {
		newSources := split(input, ',')
		config.APP.SOURCE = newSources
	}

	// Write the updated configuration
	err = writeConfig(config, fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Println("Configuration updated successfully!")
}

func split(s string, delimiter rune) []string {
	var parts []string
	start := 0
	for i, r := range s {
		if r == delimiter {
			parts = append(parts, s[start:i])
			start = i + 1
		}
	}
	parts = append(parts, s[start:])
	return parts
}
