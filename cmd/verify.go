package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/wheniwork/ultronpym/data"
)

// RootCmd handles the default commands
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the manifest file and the files list",
	Run: func(cmd *cobra.Command, args []string) {
		verifyCommand(manifestFile, signatureFile)
	},
}

func verifyCommand(manifestFile string, signatureFile string) {
	if publicKey == "" {
		out, err := exec.Command("gpg2", "--export", "--armor").Output()
		if err != nil {
			fatal(err)
		}
		publicKey = string(out)
	}
	/* Verify signature of manifest file */
	ok, err := data.VerifySignature(manifestFile, signatureFile, publicKey)
	if !ok {
		fatal(err)
	}
	success(fmt.Sprintf("%s checksum OK", manifestFile))

	/* Verify the contents of the manifest file */
	manifest := data.LoadManifest(manifestFile)
	for file, currentSha := range manifest {
		sha, err := data.ComputeChecksum(file)
		if err != nil {
			fatal(fmt.Sprintf("Error verifying: %s\n", err.Error()))
		}
		if sha != currentSha {
			fatal(fmt.Sprintf("Error checksum for file: %s", file))
		}

		success(fmt.Sprintf("Checksum OK: %s", file))
	}
	success("Valid!")
}
