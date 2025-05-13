package auth

import (
	"fmt"
	"os"
	"os/exec"
)

// Ensure user is authenticated with gcloud
func Ensure() error {
	// Check if gcloud is installed
	if _, err := exec.LookPath("gcloud"); err != nil {
		return fmt.Errorf("gcloud CLI not found in PATH: please install the Google Cloud SDK from https://cloud.google.com/sdk/docs/install")
	}
	// Check active account
	out, _ := exec.Command("gcloud", "auth", "list",
		"--filter=status:ACTIVE", "--format=value(account)").Output()
	if len(out) == 0 {
		fmt.Println("🔐  Not authenticated. Launching gcloud auth login...")
		cmd := exec.Command("gcloud", "auth", "login")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	return nil
}
