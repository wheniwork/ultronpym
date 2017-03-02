package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wheniwork/ultronpym/data"
)

// RootCmd handles the default commands
var removeCmd = &cobra.Command{
	Use:   "rm [file to remove]",
	Short: "Remove a file from the manifest",
	Run: func(cmd *cobra.Command, args []string) {
		removeCommand(manifestFile, args)
	},
}

func removeCommand(manifestFile string, files []string) {
	manifest := data.LoadManifest(manifestFile)
	for _, file := range files {
		sha, ok := manifest[file]
		if !ok {
			color.New(color.FgRed).Printf("File %s not in manifest\n", file)
			continue
		}
		yellow := color.New(color.FgYellow).SprintfFunc()
		fmt.Printf("%s: %s\n", yellow("Removed %s", file), sha)
		delete(manifest, file)
	}
	manifest.SaveManifest(manifestFile)
	success("\nFile(s) removed")
}
