package cmd

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// RootCmd handles the default commands
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Signs the manifest file and generates a signature file",
	Run: func(cmd *cobra.Command, args []string) {
		signCommand(manifestFile, signatureFile)
	},
}

func signCommand(manifestFile string, signatureFile string) {
	rehashManifest(manifestFile)
	out, err := exec.Command("gpg2", "--batch", "--yes", "--output", signatureFile, "--detach-sig", manifestFile).CombinedOutput()
	if err != nil {
		fatal(string(out))
	}
	fmt.Printf("Creating file %s\n", signatureFile)
	color.New(color.FgHiGreen).Println("🔑  Signature Created!")
}
