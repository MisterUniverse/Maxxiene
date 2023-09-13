package winfirewall

import (
	"fmt"
	"os"
	"os/exec"
)

// ExecuteCommand executes a shell command and returns its output
func ExecuteCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// ListFirewallRules lists all the firewall rules
func ListFirewallRules(ruleName string) (string, error) {
	fmt.Println("Listing firewall rules...")
	output, err := ExecuteCommand("netsh", "advfirewall", "firewall", "show", "rule", "name="+ruleName)
	return string(output), err
}

// promptUser prompts user for input with a default value
func promptUser(prompt, defaultValue string) string {
	var input string
	fmt.Print(prompt + " [" + defaultValue + "]: ")
	fmt.Scanln(&input)
	if input == "" {
		return defaultValue
	}
	return input
}

// AddFirewallRule adds a new firewall rule with user prompts
func AddFirewallRule() error {
	fmt.Println("Adding a new firewall rule.")
	ruleName := promptUser("Enter rule name. Default is:", "MyDefaultFirewallRuleName")
	protocol := promptUser("Enter protocol (TCP|UDP). Default is:", "Any")
	profile := promptUser("Enter profile (public/private/domain). Default is:", "Any")
	direction := promptUser("Enter direction (in|out). Default is:", "in")
	action := promptUser("Enter action (allow|block|bypass). Default is:", "allow")

	// Execute command with the parameters, default or user-provided
	fmt.Println("Adding firewall rule...")
	_, err := ExecuteCommand("netsh", "advfirewall", "firewall", "add", "rule",
		"name="+ruleName,
		"dir="+direction,
		"action="+action,
		"protocol="+protocol,
		"profile="+profile,
	)

	return err
}

// DeleteFirewallRule deletes an existing inbound firewall rule
func DeleteFirewallRule(ruleName string) error {
	fmt.Println("Deleting firewall rule...")
	_, err := ExecuteCommand("netsh", "advfirewall", "firewall", "delete", "rule", "name="+ruleName)
	return err
}

// EnableFirewall enables the firewall for a specific profile
func EnableFirewall(profile string) error {
	fmt.Println("Enabling firewall for profile:", profile)
	_, err := ExecuteCommand("netsh", "advfirewall", "set", profile, "state", "on")
	return err
}

// DisableFirewall disables the firewall for a specific profile
func DisableFirewall(profile string) error {
	fmt.Println("Disabling firewall for profile:", profile)
	_, err := ExecuteCommand("netsh", "advfirewall", "set", profile, "state", "off")
	return err
}

// ExportFirewallRules exports the firewall rules to a file
func ExportFirewallRules(filePath string) {
	fmt.Println("Exporting firewall rules to:", filePath)
	_, err := ExecuteCommand("netsh", "advfirewall", "export", filePath)
	if err != nil {
		fmt.Println("Failed to export firewall rules:", err)
		return
	}
	fmt.Println("Successfully exported firewall rules to:", filePath)
}

// ExportHumanReadable exports the firewall rules to a human-readable text file
func ExportHumanReadable(filePath string) {
	fmt.Println("Exporting firewall rules to human-readable format at:", filePath)
	output, err := ExecuteCommand("netsh", "advfirewall", "firewall", "show", "rule", "name=all")
	if err != nil {
		fmt.Println("Failed to fetch firewall rules:", err)
		return
	}

	err = os.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}

	fmt.Println("Successfully exported firewall rules to:", filePath)
}

// ImportFirewallRules imports the firewall rules from a file
func ImportFirewallRules(filePath string) error {
	fmt.Println("Importing firewall rules from:", filePath)
	_, err := ExecuteCommand("netsh", "advfirewall", "import", filePath)
	return err
}
