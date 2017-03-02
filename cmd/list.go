package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wheniwork/ultronpym/data"
)

// RootCmd handles the default commands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the files within the manifest",
	Run: func(cmd *cobra.Command, args []string) {
		listCommand(manifestFile)
	},
}

func listCommand(manifestFile string) {
	manifest := data.LoadManifest(manifestFile)
	for file, sha := range manifest {
		printHash(file, sha, false)
	}
}
