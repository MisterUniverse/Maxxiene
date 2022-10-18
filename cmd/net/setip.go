/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"maxx/utils"
	"strings"

	"github.com/spf13/cobra"
)

var (
	ip                 string
	mask               = "255.255.255.0"
	name               string
	elevateProcessCmds = `
	$myWindowsID=[System.Security.Principal.WindowsIdentity]::GetCurrent()
	$myWindowsPrincipal=new-object System.Security.Principal.WindowsPrincipal($myWindowsID)
	$adminRole=[System.Security.Principal.WindowsBuiltInRole]::Administrator
	$newProcess = new-object System.Diagnostics.ProcessStartInfo "PowerShell";
	$newProcess.Arguments = $MyInvocation.MyCommand.Definition.Path;
	$newProcess.Verb = "runas";
	$process = [System.Diagnostics.Process]::Start($newProcess);
	exit
`
)

// setipCmd represents the setip command
var setipCmd = &cobra.Command{
	Use:   "setip",
	Short: "Allows you to set your ip address",
	Long:  `Enter an ip address or "dhcp" to set ip to dhcp mode. Sets user to dhcp by default.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ip == "" {
			cmd.Help()
			return
		}
		psh := utils.NewShell()

		i := strings.Split(ip, ".")
		i = i[:len(i)-1]
		i = append(i, "1")

		gateway := strings.Join(i, ".")

		if strings.ToLower(ip) != "dhcp" {
			pcmd := fmt.Sprintf("netsh interface ipv4 set address name=\"%v\" static %v %v %v", name, ip, mask, gateway)
			shwConf := fmt.Sprintf("%s \n netsh interface ipv4 show config", pcmd)
			fmt.Println(shwConf)
			stdOut, stdErr, err := psh.Execute(shwConf)
			fmt.Printf("\n setip: \n StdOut : '%s' \n StdErr: '%s' \n Err: %s", strings.TrimSpace(stdOut), stdErr, err)
		} else {
			pcmd := fmt.Sprintf("netsh interface ipv4 set address name=\"%v\" source=\"dhcp\"", name)
			shwConf := fmt.Sprintf("%s \n netsh interface ipv4 show config", pcmd)
			stdOut, stdErr, err := psh.Execute(shwConf)
			fmt.Printf("\n setip: \n StdOut : '%s' \n StdErr: '%s' \n Err: %s", strings.TrimSpace(stdOut), stdErr, err)
		}
	},
}

func init() {
	setipCmd.Flags().StringVarP(&name, "name", "n", "Ethernet", "Name of network interface")
	setipCmd.Flags().StringVarP(&ip, "ipv4", "i", "", "Ip address")
	if err := setipCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}

	if err := setipCmd.MarkFlagRequired("ipv4"); err != nil {
		fmt.Println(err)
	}

	NetCmd.AddCommand(setipCmd)
}
