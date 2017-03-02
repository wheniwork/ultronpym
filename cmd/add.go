package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wheniwork/ultronpym/data"
)

// RootCmd handles the default commands
var addCmd = &cobra.Command{
	Use:   "add [files to add]",
	Short: "Add a file from the manifest",
	Run: func(cmd *cobra.Command, args []string) {
		addCommand(manifestFile, args)
	},
}

func addCommand(manifestFile string, args []string) {
	manifest := data.LoadManifest(manifestFile)
	for _, file := range args {
		sha, err := data.ComputeChecksum(file)
		if err != nil {
			color.New(color.FgRed).Printf("Error adding file: %s\n", err.Error())
			continue
		}
		oldSha, exists := manifest[file]
		printHash(file, sha, !exists || oldSha != sha)
		manifest[file] = sha
	}
	manifest.SaveManifest(manifestFile)
	success("\nFile(s) added!")
}
