package main

import (
	"fmt"
	"os"
)

const hostsFilePath = "/etc/hosts"
const backupFilePath = "/etc/hosts.bak"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: focus-tool <start|stop>")
		os.Exit(1)
	}

	command := os.Args[1]
	fmt.Println(os.Args)

	switch command {
	case "start":
		startBlocking()
	case "stop":
		stopBlocking()
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}

func startBlocking() {
	// Backup the current hosts file
	input, err := os.ReadFile(hostsFilePath)
	if err != nil {
		fmt.Println("Error reading hosts file:", err)
		os.Exit(1)
	}

	err = os.WriteFile(backupFilePath, input, 0644)
	if err != nil {
		fmt.Println("Error creating backup hosts file:", err)
		os.Exit(1)
	}

	// Add blocked URLs to hosts file
	f, err := os.OpenFile(hostsFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening hosts file:", err)
		os.Exit(1)
	}
	defer f.Close()

	for _, url := range blockedURLs {
		_, err := f.WriteString("127.0.0.1 " + url + "\n")
		if err != nil {
			fmt.Println("Error writing to hosts file:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Social media URLs blocked.")
}

func stopBlocking() {
	// Restore the hosts file from the backup
	input, err := os.ReadFile(backupFilePath)
	if err != nil {
		fmt.Println("Error reading backup hosts file:", err)
		os.Exit(1)
	}

	err = os.WriteFile(hostsFilePath, input, 0644)
	if err != nil {
		fmt.Println("Error restoring hosts file:", err)
		os.Exit(1)
	}

	fmt.Println("Social media URLs unblocked.")
}
